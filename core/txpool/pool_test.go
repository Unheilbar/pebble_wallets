package txpool

import (
	"log"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
)

var startTime = time.Unix(1697495678, 0)

func Test_txPool(t *testing.T) {
	txpool := New()

	go func() {
		<-time.After(time.Millisecond * 5)
		txes := txpool.Pending(500)
		if len(txes) < 500 {
			log.Fatal("couldn't recieve tx in time ", len(txes))
		}
		for i := 0; i < len(txes)-1; i++ {
			if txes[i].Time().UnixNano() > txes[i+1].Time().UnixNano() {
				log.Fatal("not ordered tx", txes[i].Time().UnixNano(), txes[i+1].Time().UnixNano())
			}
		}
		txes = txpool.Pending(500)
		if len(txes) < 500 {
			log.Fatal("didn't get all expected txes", len(txes))
		}
	}()
	for _, tx := range generateTxes(1000) {
		err := txpool.AddTx(tx)
		if err != nil {
			log.Fatal("err ", tx.Id().Hex())
		}
	}
}

func BenchmarkTxPoolAdd(b *testing.B) {
	txes := generateTxes(10000)
	txpool := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range rand.Perm(len(txes)) {
			txpool.AddTx(txes[v])
		}
	}
}

func BenchmarkTxPoolWithBackgroundPendingСall(b *testing.B) {
	txes := generateTxes(10000)
	txpool := New()
	b.ResetTimer()
	go func() {
		for {
			_ = txpool.Pending(1000)
			time.After(time.Millisecond * 50)
		}
	}()
	for i := 0; i < b.N; i++ {
		for _, v := range rand.Perm(len(txes)) {
			txpool.AddTx(txes[v])
		}
	}
}

func generatepooltx(i int) *types.Transaction {
	return types.NewTx(types.TxData{Id: common.BytesToHash(big.NewInt(int64(i)).Bytes())}).WithTime(startTime.Add(time.Second * time.Duration(i))) // 5 // 1697495683
}

func generateTxes(size int) []*types.Transaction {
	txs := make([]*types.Transaction, 0)
	for i := 0; i < size; i++ {
		txs = append(txs, generatepooltx(i))
	}
	return txs
}
