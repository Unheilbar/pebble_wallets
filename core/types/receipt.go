package types

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Receipt struct {
	Status uint64 `json:"status"`
	Logs   []*Log `json:"logs"              gencodec:"required"`

	// Implementation fields: These fields are added by geth when processing a transaction.
	TxHash common.Hash `json:"transactionHash" gencodec:"required"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        common.Hash    `json:"blockHash,omitempty"`
	BlockNumber      *big.Int       `json:"blockNumber,omitempty"`
	TransactionIndex uint           `json:"transactionIndex"`
	ContractAddress  common.Address `json:"contractAddress"`
}

const (
	// ReceiptStatusFailed is the status code of a transaction if execution failed.
	ReceiptStatusFailed = uint64(0)

	// ReceiptStatusSuccessful is the status code of a transaction if execution succeeded.
	ReceiptStatusSuccessful = uint64(1)
)

// Receipts implements DerivableList for receipts.
type Receipts []*Receipt

// Len returns the number of receipts in this list.
func (rs Receipts) Len() int { return len(rs) }

// EncodeIndex encodes the i'th receipt to w.
func (rs Receipts) EncodeIndex(i int, w *bytes.Buffer) {
	r := rs[i]
	data := &receiptRLP{r.Logs, r.Status}

	rlp.Encode(w, data)
}

type receiptRLP struct {
	logs   []*Log
	status uint64
}
