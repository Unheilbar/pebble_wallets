package mlsmtrie

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"github.com/celestiaorg/smt"
	"github.com/ethereum/go-ethereum/common"
)

func Test_New(t *testing.T) {
	// Initialise two new key-value store to store the nodes and values of the tree
	nodeStore := smt.NewSimpleMap()
	valueStore := smt.NewSimpleMap()
	// Initialise the tree
	tree := smt.NewSparseMerkleTree(nodeStore, valueStore, sha256.New())
	tree.Update([]byte("foo"), []byte("bar"))
	tree.Update([]byte("321"), []byte("321"))
	root := tree.Root()
	tree.Update([]byte("foo"), []byte("bar1"))
	tree.Root()

	expBar, _ := tree.Get([]byte("foo"))
	fmt.Println(string(expBar))

	tree.SetRoot(root)
	expBar2, _ := tree.Get([]byte("foo"))
	fmt.Println(string(expBar2))

	fmt.Println(common.BytesToHash(root).Hex())
}

func Test_fakeVM(t *testing.T) {
	runFakeVM()
}
