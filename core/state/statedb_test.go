package state

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

var walletsAmount = 100000
var txesAmount = 100000
var batchSize = 10000

func Test__ApplyTransactions(t *testing.T) {
	state := NewState("../../test")

	ts := newTestSuit(walletsAmount, txesAmount)
	initWallets(state, ts)
	start := time.Now()
	lp := 0
	rp := 0
	iter := 0
	for rp < len(ts.txs) {
		lp = rp
		rp += batchSize
		if rp > len(ts.txs) {
			rp = len(ts.txs)
		}

		resultHash := state.ApplyTransactions(ts.txs[lp:rp])

		iter++
		fmt.Println("iter", iter, "hash", resultHash)
	}

	res := time.Since(start)

	fmt.Println(res, float64(txesAmount)/res.Seconds(), "tx/s ")
}

func initWallets(state *StateDB, ts testSuit) {
	lp := 0
	rp := 0
	iter := 0

	for rp < len(ts.txs) {
		lp = rp
		rp += batchSize

		if rp > len(ts.txs) {
			rp = len(ts.txs)
		}

		state.InitWallets(ts.wallets[lp:rp])

		iter++
	}
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
			Value: big.NewInt(1),
			Id:    common.BytesToHash([]byte(uuid.New().String())),
		})
	}
	return txes
}

func generateAccounts(size int) []common.Address {
	var s []common.Address
	for i := 0; i < size; i++ {
		k := common.BytesToAddress(crypto.Keccak256(big.NewInt(int64(i)).Bytes())[12:])

		s = append(s, k)
	}
	return s
}
