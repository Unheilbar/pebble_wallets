package core

import (
	"github.com/Unheilbar/pebbke_wallets/core/types"
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
}

type Processor interface {
	Process(block types.Block, statedb *state.StateDB) []*types.Receipt
}
