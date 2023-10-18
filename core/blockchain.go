package core

import (
	"fmt"
	"log"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Unheilbar/pebbke_wallets/binding"
	"github.com/Unheilbar/pebbke_wallets/core/rawdb"
	"github.com/Unheilbar/pebbke_wallets/core/state"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/Unheilbar/pebbke_wallets/core/vm"
	"github.com/Unheilbar/pebbke_wallets/trie"
	"github.com/Unheilbar/pebbke_wallets/trie/triedb/hashdb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
)

type Blockchain struct {
	// Databases
	db           ethdb.Database
	triedb       *trie.Database              // The database handler for maintaining trie nodes.
	stateCache   state.Database              // State database to reuse between imports (contains state cache)
	processor    Processor                   // Block transaction processor interface
	currentBlock atomic.Pointer[types.Block] // Current head of the chain

	chainmu sync.RWMutex // blockchain insertion lock

	prefetcher Prefetcher //
}

type Processor interface {
	Process(block types.Block, statedb *state.StateDB) []*types.Receipt
}

// Prefetcher is an interface for pre-caching transaction signatures and state.
type Prefetcher interface {
	// Prefetch processes the state changes according to the Ethereum rules by running
	// the transaction messages using the statedb, but any changes are discarded. The
	// only goal is to pre-cache transaction signatures and state trie nodes.
	Prefetch(block *types.Block, statedb *state.StateDB, cfg vm.Config, interrupt *atomic.Bool)
}

type CacheConfig struct {
	TrieCleanLimit      int           // Memory allowance (MB) to use for caching trie nodes in memory
	TrieCleanNoPrefetch bool          // Whether to disable heuristic state prefetching for followup blocks
	TrieDirtyLimit      int           // Memory limit (MB) at which to start flushing dirty trie nodes to disk
	TrieDirtyDisabled   bool          // Whether to disable trie write caching and GC altogether (archive node)
	TrieTimeLimit       time.Duration // Time limit after which to flush the current in-memory trie to disk
	SnapshotLimit       int           // Memory allowance (MB) to use for caching snapshot entries in memory
	Preimages           bool          // Whether to store preimage of trie key to the disk
	StateHistory        uint64        // Number of blocks from head whose state histories are reserved.
	StateScheme         string        // Scheme used to store ethereum states and merkle tree nodes on top

	SnapshotNoBuild bool // Whether the background generation is allowed
	SnapshotWait    bool // Wait for snapshot construction on startup. TODO(karalabe): This is a dirty hack for testing, nuke it
}

// defaultCacheConfig are the default caching values if none are specified by the
// user (also used during testing).
var defaultCacheConfig = &CacheConfig{
	TrieCleanLimit: 256,
	TrieDirtyLimit: 256,
	TrieTimeLimit:  5 * time.Minute,
	SnapshotLimit:  256,
	SnapshotWait:   true,
	StateScheme:    rawdb.HashScheme,
}

// triedbConfig derives the configures for trie database.
func (c *CacheConfig) triedbConfig() *trie.Config {
	config := &trie.Config{Preimages: c.Preimages}
	if c.StateScheme == rawdb.HashScheme {
		config.HashDB = &hashdb.Config{
			CleanCacheSize: c.TrieCleanLimit * 1024 * 1024,
		}
	}

	return config
}

func NewBlockchain(rdb ethdb.Database) *Blockchain {
	bc := &Blockchain{
		db: rdb,
	}
	bc.chainmu.Lock()
	defer bc.chainmu.Unlock()
	// Open trie database with provided config
	bc.triedb = trie.NewDatabase(rdb, defaultCacheConfig.triedbConfig())
	bc.stateCache = state.NewDatabaseWithNodeDB(bc.db, bc.triedb)

	// TODO too dirty, just for test reasons
	var blockID int64 = 1
	deployStartTime := time.Now()

	block := types.NewBlockWithHeader(&types.Header{
		Number: big.NewInt(blockID),
	})
	block.Transactions = []*types.Transaction{getContractDeployTX(common.Hex2Bytes(binding.StorageMetaData.Bin[2:]))}
	statedb, err := state.New(common.Hash{}, bc.stateCache, nil)
	stProcessor := NewStateProcessor()
	receipts := stProcessor.Process(*block, statedb)
	receipt := receipts[0]
	contrAddr := receipt.ContractAddress

	newRoot, err := statedb.Commit(uint64(blockID), false)
	if err != nil {
		log.Fatal("err commit sb deploy", err)
	}

	block.Header().Root = newRoot
	bc.currentBlock.Store(block)
	err = bc.stateCache.TrieDB().Commit(newRoot, false)
	if err != nil {
		log.Fatal("err commit tdb deploy", err)
	}
	sb, err := bc.StateAt(newRoot)
	if err != nil {
		panic(err)
	}
	code := sb.GetCode(contrAddr)
	if code == nil {
		panic("err create blockchain")
	}
	log.Println("test storage addr", contrAddr, receipt.Status, "root", newRoot, "time", time.Since(deployStartTime))
	return bc
}

// StateAt returns a new mutable public state and a new mutable private state repo
// based on a particular point in time. The returned private state repo can be used
// to obtain a mutable private state for a given PSI
func (bc *Blockchain) StateAt(root common.Hash) (*state.StateDB, error) {
	publicStateDb, publicStateDbErr := state.New(root, bc.stateCache, nil)

	return publicStateDb, publicStateDbErr
}

// RAFT ONLY
func (bc *Blockchain) CommitBlockWithState(blockNumber uint64, state *state.StateDB) error {
	// check if consensus is not Raft

	bc.chainmu.Lock()
	defer bc.chainmu.Unlock()
	check := state.GetCode(common.HexToAddress("0x6027946B05e7ab6Ef245093622AB18eaD5453877"))
	if check == nil {
		panic("precheck failed")
	}
	root, err := state.Commit(blockNumber, true)
	if err != nil {
		return fmt.Errorf("error committing public state: %v", err)
	}

	state, err = bc.StateAt(root)
	if err != nil {
		panic(err)
	}
	check = state.GetCode(common.HexToAddress("0x6027946B05e7ab6Ef245093622AB18eaD5453877"))
	if check == nil {
		panic("check failed")
	}
	return nil
}

func getSpeed(txes int, interval time.Duration) float64 {
	return float64(txes) / interval.Seconds()
}

func (bc *Blockchain) InsertChain(block *types.Block) error {
	bc.chainmu.Lock()
	defer bc.chainmu.Unlock()
	elapsed := time.Since(time.Unix(0, int64(block.Time())))
	err := bc.writeBlockAndSetHead(block)
	log.Println("🔨  Insert chain block", "number", block.Number(), "hash", fmt.Sprintf("%x", block.Hash().Bytes()[:4]), "elapsed", elapsed.Seconds(), "len(txs): ", len(block.Transactions), getSpeed(len(block.Transactions), elapsed), "tx/s ")
	return err
}

func (bc *Blockchain) writeBlockWithState(block *types.Block, state *state.StateDB) error {

	// Irrelevant of the canonical status, write the block itself to the database.
	//
	// Note all the components of block(td, hash->number map, header, body, receipts)
	// should be written atomically. BlockBatch is used for containing all components.
	blockBatch := bc.db.NewBatch()
	rlpBlock, err := rlp.EncodeToBytes(block)
	if err != nil {
		return err
	}
	rawdb.WriteBodyRLP(blockBatch, block.Hash(), block.Number().Uint64(), rlpBlock)

	if err := blockBatch.Write(); err != nil {
		log.Fatal("Failed to write block into disk", "err", err)
	}
	// Commit all cached state changes into underlying memory database.
	root, err := state.Commit(block.NumberU64(), true)
	if err != nil {
		return err
	}

	return bc.stateCache.TrieDB().Commit(root, false)
}

// writeBlockAndSetHead is the internal implementation of WriteBlockAndSetHead.
// This function expects the chain mutex to be held.
func (bc *Blockchain) writeBlockAndSetHead(block *types.Block) error {
	state, err := bc.StateAt(block.Root())
	if err != nil {
		return err
	}

	if err := bc.writeBlockWithState(block, state); err != nil {
		return err
	}

	return nil
}

func (bc *Blockchain) writeHeadBlock(block *types.Block) {
	// Add the block to the canonical chain number scheme and mark as the head
	batch := bc.db.NewBatch()
	rawdb.WriteHeadHeaderHash(batch, block.Hash())
	rawdb.WriteHeadFastBlockHash(batch, block.Hash())
	rawdb.WriteCanonicalHash(batch, block.Hash(), block.NumberU64())
	rawdb.WriteTxLookupEntriesByBlock(batch, block)
	rawdb.WriteHeadBlockHash(batch, block.Hash())

	// Flush the whole batch into the disk, exit the node if failed
	if err := batch.Write(); err != nil {
		log.Fatal("Failed to update chain indexes and markers", "err", err)
	}
}

func (bc *Blockchain) CurrentBlock() *types.Block {
	return bc.currentBlock.Load()
}

// only for stess minter tests PEBBLE REMOVE LATER
func getContractDeployTX(contrCode []byte) *types.Transaction {
	return types.NewTx(types.TxData{
		From: common.Address{},
		Data: contrCode,
	})
}
