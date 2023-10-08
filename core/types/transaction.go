package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	From      common.Address
	To        common.Address
	Id        common.Hash
	Signature common.Hash
	Value     *big.Int
}
