package core

import (
	"fmt"
	"sync"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/trie"
)

type Blockchain struct {
	// Databases
	db         ethdb.Database
	triedb     *trie.Database // The database handler for maintaining trie nodes.
	stateCache state.Database // State database to reuse between imports (contains state cache)
	processor  Processor      // Block transaction processor interface

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
