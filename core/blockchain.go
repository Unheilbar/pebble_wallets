package core

import (
	"fmt"
	"log"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Unheilbar/pebbke_wallets/binding"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

type Blockchain struct {
	// Databases
	db           ethdb.Database
	triedb       *trie.Database              // The database handler for maintaining trie nodes.
	stateCache   state.Database              // State database to reuse between imports (contains state cache)
	processor    Processor                   // Block transaction processor interface
	currentBlock atomic.Pointer[types.Block] // Current head of the chain

	chainmu sync.RWMutex // blockchain insertion lock
}

type Processor interface {
	Process(block types.Block, statedb *state.StateDB) []*types.Receipt
}

func NewBlockchain(rdb ethdb.Database) *Blockchain {
	bc := &Blockchain{
		db: rdb,
	}
	// Open trie database with provided config
	bc.triedb = trie.NewDatabase(rdb, trie.HashDefaults)
	bc.stateCache = state.NewDatabaseWithNodeDB(bc.db, bc.triedb)

	// TODO too dirty, just for test reasons
	var blockID int64 = 1
	deployStartTime := time.Now()

	block := types.Block{
		Header: &types.Header{
			Number: big.NewInt(blockID),
		},
		Transactions: []*types.Transaction{getContractDeployTX(common.Hex2Bytes(binding.StorageMetaData.Bin[2:]))},
	}
	statedb, err := state.New(common.Hash{}, bc.stateCache, nil)
	stProcessor := NewStateProcessor(getDefaultCfg())
	receipts := stProcessor.Process(block, statedb)
	receipt := receipts[0]
	contrAddr := receipt.ContractAddress

	newRoot, err := statedb.Commit(uint64(blockID), true)
	if err != nil {
		log.Fatal("err commit sb deploy", err)
	}
	block.Header.Root = newRoot
	bc.currentBlock.Store(&block)
	err = bc.stateCache.TrieDB().Commit(newRoot, false)
	if err != nil {
		log.Fatal("err commit tdb deploy", err)
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
	if _, err := state.Commit(blockNumber, true); err != nil {
		return fmt.Errorf("error committing public state: %v", err)
	}

	return nil
}

func (bc *Blockchain) InsertChain(block *types.Block) error {
	bc.chainmu.Lock()
	defer bc.chainmu.Unlock()
	return bc.writeBlockAndSetHead(block)

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
	// rawdb.WriteTxLookupEntriesByBlock(batch, block) TODO cant write lookup because no compatibility
	rawdb.WriteHeadBlockHash(batch, block.Hash())

	// Flush the whole batch into the disk, exit the node if failed
	if err := batch.Write(); err != nil {
		log.Fatal("Failed to update chain indexes and markers", "err", err)
	}
}

func (bc *Blockchain) CurrentBlock() *types.Block {
	return bc.currentBlock.Load()
}

func getContractDeployTX(contrCode []byte) *types.Transaction {
	return &types.Transaction{
		From:  common.Address{},
		To:    common.Address{},
		Input: contrCode,
	}
}
