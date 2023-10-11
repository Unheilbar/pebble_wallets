package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Block struct {
	Transactions []*Transaction
	Number       *big.Int
	Hash         common.Hash
}
