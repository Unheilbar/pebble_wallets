package types

import (
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/trie"
	"golang.org/x/crypto/sha3"
)

type Header struct {
	ParentHash  common.Hash
	TxHash      common.Hash
	ReceiptHash common.Hash
	Number      *big.Int
	Root        common.Hash // state root
}

type Block struct {
	Header *Header

	Transactions []*Transaction
	Receipts     []*Receipt
}

func (block *Block) Hash() common.Hash {
	return rlpHash(block.Number)
}

var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

func NewBlock(header *Header, txs []*Transaction, receipts []*Receipt) *Block {
	b := &Block{
		Transactions: txs,
		Receipts:     receipts,
	}

	header.ReceiptHash = DeriveSha(Receipts(receipts), trie.NewStackTrie(nil))
	header.TxHash = DeriveSha(Transactions(txs), trie.NewStackTrie(nil))

	b.Header = header
	return b
}
func (b *Block) Number() *big.Int  { return new(big.Int).Set(b.Header.Number) }
func (b *Block) Root() common.Hash { return b.Header.Root }
