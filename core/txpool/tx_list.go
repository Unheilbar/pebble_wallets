package txpool

import (
	"github.com/Unheilbar/pebble_wallets/core/types"
)

// An Item is something we manage in a priority queue.
type timestampTxItem struct {
	tx       *types.Transaction // The value of the item; arbitrary.
	priority int64              // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type timestampTxList []*timestampTxItem

func (pq timestampTxList) Len() int { return len(pq) }

func (pq timestampTxList) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq timestampTxList) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *timestampTxList) Push(x any) {
	n := len(*pq)
	item := x.(*timestampTxItem)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *timestampTxList) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
