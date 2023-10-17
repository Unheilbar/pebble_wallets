package txpool

import (
	"container/heap"
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
)

func Test__timestampTxList(t *testing.T) {
	// Some items and their priorities.
	startTime := time.Unix(1697495678, 0)
	tx3 := types.NewTx(types.TxData{Id: common.HexToHash("0x1")}).WithTime(startTime.Add(time.Second * 5)) // 5 // 1697495683
	tx1 := types.NewTx(types.TxData{Id: common.HexToHash("0x2")}).WithTime(startTime.Add(time.Second))     // 1 // 1697495679
	tx2 := types.NewTx(types.TxData{Id: common.HexToHash("0x3")}).WithTime(startTime.Add(time.Second * 3)) // 2 // 1697495681
	tx5 := types.NewTx(types.TxData{Id: common.HexToHash("0x3")}).WithTime(startTime.Add(time.Second * 3)) // 3 // 1697495681
	tx4 := types.NewTx(types.TxData{Id: common.HexToHash("0x4")}).WithTime(startTime.Add(time.Second * 4)) // 4 //  1697495682
	unorderedTxs := []*types.Transaction{
		tx1,
		tx2,
		tx3,
	}

	expectedTxs := []*types.Transaction{tx1, tx2, tx5, tx4, tx3}

	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(timestampTxList, len(unorderedTxs))
	i := 0
	for _, utx := range unorderedTxs {
		pq[i] = &timestampTxItem{
			tx:       utx,
			priority: utx.Time().UnixMicro(),
			index:    i,
		}
		i++
	}

	heap.Init(&pq)

	// add one with not replacing time
	heap.Push(&pq, &timestampTxItem{tx: tx4, priority: tx4.Time().UnixMicro()})
	heap.Push(&pq, &timestampTxItem{tx: tx5, priority: tx5.Time().UnixMicro()})

	// Take the items out; they arrive in decreasing priority order.
	ii := 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*timestampTxItem)
		if expectedTxs[ii].Id() != item.tx.Id() {
			log.Fatal(fmt.Sprintf("got Id %s instead of %s ", expectedTxs[ii].Id().Hex(), item.tx.Id().Hex()))
		}
		ii++
		// fmt.Println(item.tx.Time().Unix(), item.tx.Id())
	}
}

func BenchmarkListAdd(b *testing.B) {
	txs := make([]*types.Transaction, 3000)
	for i := 0; i < len(txs); i++ {
		txs[i] = transaction(int64(i), common.BytesToHash([]byte(fmt.Sprint(i))))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list := new(timestampTxList)
		for _, v := range rand.Perm(len(txs)) {
			list.Push(timestamptxitem(txs[v]))
		}
	}
}

func timestamptxitem(t *types.Transaction) *timestampTxItem {
	return &timestampTxItem{tx: t, priority: t.Time().Unix()}
}

func transaction(timestamp int64, id common.Hash) *types.Transaction {
	return types.NewTx(types.TxData{Id: id}).WithTime(time.Unix(timestamp, 0))
}
