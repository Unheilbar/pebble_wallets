package types

import (
	"bytes"
	"sync/atomic"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type TxData struct {
	From      common.Address
	To        *common.Address
	Id        common.Hash
	Signature []byte
	Data      []byte

	AccessList // TODO for later
}

type Transaction struct {
	inner TxData    // Consensus contents of a transaction
	time  time.Time // Time first seen locally (spam avoidance)

	// caches
	hash atomic.Value
}

func (tx *Transaction) Signature() []byte {
	return tx.inner.Signature
}

func (tx *Transaction) Data() []byte {
	return tx.inner.Data
}

func (tx *Transaction) Hash() common.Hash {
	if hash := tx.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}

	var h common.Hash

	h = rlpHash(tx.inner)

	tx.hash.Store(h)
	return h
}

// NewTx creates a new transaction.
func NewTx(inner TxData) *Transaction {
	tx := new(Transaction)
	tx.setDecoded(inner.copy())
	return tx
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *TxData) copy() TxData {
	cpy := TxData{
		To:        copyAddressPtr(tx.To),
		Data:      common.CopyBytes(tx.Data),
		Signature: common.CopyBytes(tx.Signature),
		Id:        tx.Id,
		From:      tx.From,
	}

	return cpy
}

func (tx *Transaction) WithSignature(sig []byte) *Transaction {
	tx.inner.Signature = sig
	return tx
}

func (tx *Transaction) WithTime(t time.Time) *Transaction {
	tx.time = t
	return tx
}

// copyAddressPtr copies an address.
func copyAddressPtr(a *common.Address) *common.Address {
	if a == nil {
		return nil
	}
	cpy := *a
	return &cpy
}

// setDecoded sets the inner transaction and size after decoding.
func (tx *Transaction) setDecoded(inner TxData) {
	tx.inner = inner
	tx.time = time.Now()
}

// Transactions implements DerivableList for transactions.
type Transactions []*Transaction

func (tx *Transaction) From() common.Address {
	return tx.inner.From
}

func (tx *Transaction) To() *common.Address {
	return tx.inner.To
}

func (tx *Transaction) Time() time.Time {
	return tx.time
}

func (tx *Transaction) Id() common.Hash {
	return tx.inner.Id
}

// Len returns the length of s.
func (s Transactions) Len() int { return len(s) }

// EncodeIndex encodes the i'th transaction to w. Note that this does not check for errors
// because we assume that *Transaction will only ever contain valid txs that were either
// constructed by decoding or via public API in this package.
func (s Transactions) EncodeIndex(i int, w *bytes.Buffer) {
	tx := s[i].inner

	rlp.Encode(w, tx)
}

// AccessList is an EIP-2930 access list.
type AccessList []AccessTuple

// AccessTuple is the element type of an access list.
type AccessTuple struct {
	Address     common.Address `json:"address"     gencodec:"required"`
	StorageKeys []common.Hash  `json:"storageKeys" gencodec:"required"`
}
