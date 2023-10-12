package core

import "github.com/Unheilbar/pebbke_wallets/core/types"

type NewTxsEvent struct{ Txs []*types.Transaction }
