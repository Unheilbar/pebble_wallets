package state

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var walletsAmount = 100000
var txesAmount = 100000

func Test__ApplyTransactions(t *testing.T) {
	state := NewState("../../test")

	ts := newTestSuit(walletsAmount, txesAmount)

	state.InitWallets(ts.wallets)

	start := time.Now()

	state.ApplyTransactions(ts.txs)
	res := time.Since(start)
	fmt.Println(res, float64(txesAmount)/res.Seconds(), "tx/s")
}

type testSuit struct {
	wallets []common.Address
	txs     []*types.Transaction
}

func newTestSuit(accs int, txes int) testSuit {
	var ts testSuit
	ts.wallets = generateAccounts(accs)
	ts.txs = generateTxes(ts.wallets, txes)
	return ts
}

func generateTxes(addrs []common.Address, size int) []*types.Transaction {
	var txes []*types.Transaction
	var accPtr int
	for i := 0; i < size; i++ {
		fromIdx := accPtr % len(addrs)
		accPtr += 1
		toIdx := accPtr % len(addrs)
		accPtr += 1
		txes = append(txes, &types.Transaction{
			From:  addrs[fromIdx],
			To:    addrs[toIdx],
			Value: 1,
		})
	}
	return txes
}

func generateAccounts(size int) []common.Address {
	var s []common.Address
	for i := 0; i < size; i++ {
		k, err := crypto.GenerateKey()
		if err != nil {
			log.Fatal(err)
		}
		s = append(s, crypto.PubkeyToAddress(k.PublicKey))
	}
	return s
}
