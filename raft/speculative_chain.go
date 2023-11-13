package raft

import (
	"github.com/Unheilbar/pebble_wallets/core/types"
	"github.com/oleiade/lane"
)

type speculativeChain struct {
	head            *types.Block
	unappliedBlocks *lane.Deque
}

// Append a new speculative block
func (chain *speculativeChain) extend(block *types.Block) {
	chain.head = block
	// chain.recordProposedTransactions(block.Transactions()) // TODO so we won't take txes from tx pool when they come there again
	chain.unappliedBlocks.Append(block)
}

func (chain *speculativeChain) clear(block *types.Block) {
	chain.head = block
}
