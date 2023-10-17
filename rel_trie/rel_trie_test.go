package rel_trie

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func Test__GetState(t *testing.T) {
	cases := []struct {
		storageRoot common.Hash
		key         common.Hash
		val         common.Hash
	}{
		{
			common.HexToHash("0x1"),
			common.HexToHash("0x2"),
			common.HexToHash("0x3"),
		},
	}

	for i, tt := range cases {
		fmt.Println("case i")
	}
}
