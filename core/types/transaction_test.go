package types

import (
	"log"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

func Test_EncodeTransaction(t *testing.T) {
	to := common.HexToAddress("0xb84A87293A2f2A9387DcE04145bB5d97942c1129")
	tx := NewTx(TxData{
		From: common.HexToAddress("0xb84A87293A2f2A9387DcE04145bB5d97942c1129"),
		To:   &to,
	})
	tt := time.Now()
	tx.WithTime(tt)

	var btx = &Transaction{}

	b, err := rlp.EncodeToBytes(tx)
	if err != nil {
		log.Fatal(err)
	}

	err = rlp.DecodeBytes(b, &btx)
	if err != nil {
		log.Fatal(err)
	}
}
