package core

import "github.com/Unheilbar/pebble_wallets/core/types"

type NewTxsEvent struct{ Txs []*types.Transaction }

type ChainHeadEvent struct{ Block *types.Block }
