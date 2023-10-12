package types

import (
	"bytes"
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type Receipt struct {
	Status       uint64 `json:"status"`
	Logs         []*Log `json:"logs"              gencodec:"required"`
	RevertReason []byte `json:"revertReason"              gencodec:"required"`

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

// ReceiptForStorage is a wrapper around a Receipt with RLP serialization
// that omits the Bloom field and deserialization that re-computes it.
type ReceiptForStorage Receipt

// EncodeRLP implements rlp.Encoder, and flattens all content fields of a receipt
// into an RLP stream.
func (r *ReceiptForStorage) EncodeRLP(_w io.Writer) error {
	w := rlp.NewEncoderBuffer(_w)
	outerList := w.List()

	logList := w.List()
	for _, log := range r.Logs {
		if err := log.EncodeRLP(w); err != nil {
			return err
		}
	}
	w.ListEnd(logList)
	w.ListEnd(outerList)
	return w.Flush()
}

var (
	receiptStatusFailedRLP     = []byte{}
	receiptStatusSuccessfulRLP = []byte{0x01}
)

func (r *Receipt) setStatus(postStateOrStatus []byte) error {
	switch {
	case bytes.Equal(postStateOrStatus, receiptStatusSuccessfulRLP):
		r.Status = ReceiptStatusSuccessful
	case bytes.Equal(postStateOrStatus, receiptStatusFailedRLP):
		r.Status = ReceiptStatusFailed
	default:
		return fmt.Errorf("invalid receipt status %x", postStateOrStatus)
	}
	return nil
}

type storedReceiptRLP struct {
	Logs []*Log
}

// DecodeRLP implements rlp.Decoder, and loads both consensus and implementation
// fields of a receipt from an RLP stream.
func (r *ReceiptForStorage) DecodeRLP(s *rlp.Stream) error {
	var stored storedReceiptRLP
	if err := s.Decode(&stored); err != nil {
		return err
	}

	r.Logs = stored.Logs

	return nil
}

func (obj *Log) EncodeRLP(_w io.Writer) error {
	w := rlp.NewEncoderBuffer(_w)
	_tmp0 := w.List()
	w.WriteBytes(obj.Data)
	w.ListEnd(_tmp0)
	return w.Flush()
}
