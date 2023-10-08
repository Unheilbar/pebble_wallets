package types

import "github.com/ethereum/go-ethereum/common"

type Transaction struct {
	From  common.Address
	To    common.Address
	Id    common.Hash
	Value uint
}
