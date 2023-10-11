package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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
