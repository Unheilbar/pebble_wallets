package types

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type Transaction struct {
	From       common.Address
	To         common.Address
	Id         common.Hash
	Signature  []byte
	Input      []byte
	AccessList // TODO for later
}

func (tx *Transaction) Hash() common.Hash {
	val, _ := rlp.EncodeToBytes(struct {
		From  common.Address
		To    common.Address
		Id    common.Hash
		Input []byte
	}{
		tx.From,
		tx.To,
		tx.Id,
		tx.Input,
	})
	return crypto.Keccak256Hash(val)
}

// Transactions implements DerivableList for transactions.
type Transactions []*Transaction

// Len returns the length of s.
func (s Transactions) Len() int { return len(s) }

// EncodeIndex encodes the i'th transaction to w. Note that this does not check for errors
// because we assume that *Transaction will only ever contain valid txs that were either
// constructed by decoding or via public API in this package.
func (s Transactions) EncodeIndex(i int, w *bytes.Buffer) {
	tx := s[i]

	rlp.Encode(w, tx)
}

// AccessList is an EIP-2930 access list.
type AccessList []AccessTuple

// AccessTuple is the element type of an access list.
type AccessTuple struct {
	Address     common.Address `json:"address"     gencodec:"required"`
	StorageKeys []common.Hash  `json:"storageKeys" gencodec:"required"`
}
