package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type Transaction struct {
	From      common.Address
	To        common.Address
	Id        common.Hash
	Signature []byte
	Input     []byte
	Value     *big.Int
}

func (tx *Transaction) Hash() common.Hash {
	val, _ := rlp.EncodeToBytes(struct {
		From  common.Address
		To    common.Address
		Id    common.Hash
		Input []byte
		Value *big.Int
	}{
		tx.From,
		tx.To,
		tx.Id,
		tx.Input,
		tx.Value,
	})
	return crypto.Keccak256Hash(val)
}
