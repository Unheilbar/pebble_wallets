package raft

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"log"

	"github.com/Unheilbar/pebble_wallets/core"
	"github.com/Unheilbar/pebble_wallets/core/state"
	"github.com/Unheilbar/pebble_wallets/core/types"
	"github.com/Unheilbar/pebble_wallets/trie"
	"github.com/eapache/channels"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
	"github.com/oleiade/lane"
)

const maxBlockSize = 300

type minter struct {
	mu               sync.Mutex
	mux              *event.TypeMux
	chain            *core.Blockchain
	chainDb          ethdb.Database
	minting          int32 // Atomic status counter
	shouldMine       *channels.RingChannel
	blockTime        time.Duration
	speculativeChain *speculativeChain
	eth              *RaftService
	// invalidRaftOrderingChan chan InvalidRaftOrdering
	chainHeadSub  event.Subscription
	txPreChan     chan core.NewTxsEvent
	proposeBlockC chan *types.Block
}

func newMinter(eth *RaftService, blockTime time.Duration, proposeBlockC chan *types.Block) *minter {
	minter := &minter{
		eth:              eth,
		chainDb:          eth.ChainDb(),
		chain:            eth.Blockchain(),
		shouldMine:       channels.NewRingChannel(1),
		blockTime:        blockTime,
		speculativeChain: newSpeculativeChain(),

		txPreChan:     make(chan core.NewTxsEvent, 4096),
		proposeBlockC: proposeBlockC,
	}

	minter.speculativeChain.clear(minter.chain.CurrentBlock())
	// go minter.eventLoop() // we don't catch events
	go minter.mintingLoop()
	go minter.eventLoop()

	return minter
}

func (minter *minter) start() {
	atomic.StoreInt32(&minter.minting, 1)
	minter.requestMinting()
}

// Notify the minting loop that minting should occur, if it's not already been
// requested. Due to the use of a RingChannel, this function is idempotent if
// called multiple times before the minting occurs.
func (minter *minter) requestMinting() {
	minter.shouldMine.In() <- struct{}{}
}

func newSpeculativeChain() *speculativeChain {
	return &speculativeChain{
		head:            nil,
		unappliedBlocks: lane.NewDeque(),
	}
}

type work struct {
	publicState *state.StateDB
	Block       *types.Block
	header      *types.Header
}

// Assumes mu is held.
func (minter *minter) createWork() *work {
	parent := minter.speculativeChain.head
	parentNumber := parent.Number()
	tstamp := generateNanoTimestamp(parent)

	newBlockNumber := parentNumber.Add(parentNumber, common.Big1)
	header := &types.Header{
		ParentHash: parent.Hash(),
		Number:     newBlockNumber,
		Time:       uint64(tstamp),
	}

	publicState, err := minter.chain.StateAt(parent.Root())
	if err != nil {
		panic(fmt.Sprint("failed to get parent state: ", err))
	}
	fmt.Println("minter blockNumber", newBlockNumber, "apply on top", parent.Root())
	return &work{
		publicState: publicState,
		header:      header,
	}
}

func (minter *minter) mintNewBlock() {
	minter.mu.Lock()
	defer minter.mu.Unlock()
	transactions := minter.getTransactions()
	if len(transactions) == 0 {
		return
	}
	work := minter.createWork()

	committedTxes, receipts := work.commitTransactions(transactions, minter.chain)
	txCount := len(committedTxes)

	if txCount == 0 {
		// log.Print("Not minting a new block since there are no pending transactions")
		return
	}

	header := work.header

	header.Root = work.publicState.IntermediateRoot(true)

	// update block hash since it is now available, but was not when the
	// receipt/log of individual transactions were created:
	headerHash := header.Hash()
	var logs int
	for _, r := range receipts {
		if r.Logs != nil {
			logs++
			for _, l := range r.Logs {
				l.BlockHash = headerHash
				l.TxHash = r.TxHash
			}
		}
	}

	block := types.NewBlock(header, committedTxes, receipts, trie.NewStackTrie(nil))

	log.Println("Generated next block", "block num", block.Number(), " num txes ", txCount, "len logs", logs, "root", header.Root)

	if err := minter.chain.CommitBlockWithState(header.Number.Uint64(), work.publicState); err != nil {
		panic(err)
	}

	minter.speculativeChain.extend(block)

	// minter.mux.Post(core.NewMinedBlockEvent{Block: block})

	elapsed := time.Since(time.Unix(0, int64(header.Time)))
	log.Println("🔨  Mined block", "number", block.Number(), "hash", fmt.Sprintf("%x", block.Hash().Bytes()[:4]), " elapsed ", elapsed, " len(txs): ", len(committedTxes))

	//propose block
	minter.proposeBlockC <- block
}

func (env *work) commitTransactions(txs []*types.Transaction, bc *core.Blockchain) (types.Transactions, types.Receipts) {
	// commit transactions
	txes, receipts, _ := core.ApplyTransactions(bc, env.publicState, env.header, txs) // TODO probably won't be any errors here, since we dont control signal interrupt yet

	return txes, receipts
}

func (minter *minter) getTransactions() []*types.Transaction {
	return minter.eth.TxPool().Pending(maxBlockSize)
}

func (minter *minter) mintingLoop() {
	throttledMintNewBlock := throttle(minter.blockTime, func() {
		if atomic.LoadInt32(&minter.minting) == 1 {
			minter.mintNewBlock()
		} else {
			log.Print("im not a minter")
		}
	})

	for range minter.shouldMine.Out() {
		throttledMintNewBlock()
	}
}

func (minter *minter) eventLoop() {
	for {
		select {
		case <-minter.txPreChan:
			if atomic.LoadInt32(&minter.minting) == 1 {
				minter.requestMinting()
			}
		case <-time.After(time.Millisecond * 50):
			if atomic.LoadInt32(&minter.minting) == 1 {
				minter.requestMinting()
			}
		}
	}

}

func generateNanoTimestamp(parent *types.Block) (tstamp int64) {
	parentTime := int64(parent.Time())
	tstamp = time.Now().UnixNano()

	if parentTime >= tstamp {
		// Each successive block needs to be after its predecessor.
		tstamp = parentTime + 1
	}

	return
}

// Returns a wrapper around no-arg func `f` which can be called without limit
// and returns immediately: this will call the underlying func `f` at most once
// every `rate`. If this function is called more than once before the underlying
// `f` is invoked (per this rate limiting), `f` will only be called *once*.
//
// TODO(joel): this has a small bug in that you can't call it *immediately* when
// first allocated.
func throttle(rate time.Duration, f func()) func() {
	request := channels.NewRingChannel(1)

	// every tick, block waiting for another request. then serve it immediately
	go func() {
		ticker := time.NewTicker(rate)
		defer ticker.Stop()

		for range ticker.C {
			<-request.Out()
			f()
		}
	}()

	return func() {
		request.In() <- struct{}{}
	}
}
