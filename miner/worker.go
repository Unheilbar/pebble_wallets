package miner

import (
	"sync"
	"sync/atomic"
	"time"

	"github.com/Unheilbar/pebbke_wallets/core"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/event"
)

// environment is the worker's current environment and holds all
// information of the sealing block generation.
type environment struct {
	// signer   types.Signer
	state    *state.StateDB // apply state changes here
	tcount   int            // tx count in cycle
	coinbase common.Address

	header   *types.Header
	txs      []*types.Transaction
	receipts []*types.Receipt
}

type worker struct {
	//	engine      consensus.Engine TODO add engine

	chain *core.Blockchain
	txsCh chan core.NewTxsEvent
	// Feeds
	pendingLogsFeed event.Feed

	// Channels

	resultCh           chan *types.Block
	startCh            chan struct{}
	exitCh             chan struct{}
	resubmitIntervalCh chan time.Duration

	wg sync.WaitGroup

	coinbase common.Address
	extra    []byte

	pendingMu sync.RWMutex

	// External functions
	isLocalBlock func(header *types.Header) bool // Function used to determine whether the specified block is mined by local miner.
}

func (w *worker) prepareWork() (*environment, error) {
	return nil, nil //TODO setup engine, current environment from blockchain, new header etc.
}

func (w *worker) commitTransactions(env *environment, txs []*types.Transaction, interrupt *atomic.Int32) error {
	// commit transactions
	receipts, err := core.ApplyTransactions(w.chain, env.state, env.header, txs) // TODO probably won't be any errors here, since we dont control signal interrupt yet
	if err != nil {
		return err
	}
	env.txs = append(env.txs, txs...)
	env.receipts = append(env.receipts, receipts...)
	return nil
}

func (w *worker) mainLoop() {

}
