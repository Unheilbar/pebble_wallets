package tooltests

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const genesisPrivate = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"

func Test_Some(t *testing.T) {
	privateKey, err := crypto.HexToECDSA(genesisPrivate)
	if err != nil {
		log.Fatal(err)
	}
	from := common.BytesToAddress(crypto.Keccak256(crypto.FromECDSAPub(&privateKey.PublicKey)[1:])[12:])

	tx := &types.Transaction{
		From:  from,
		To:    common.HexToAddress("1"),
		Input: []byte{1, 2, 3},
	}
	signature, err := crypto.Sign(tx.Hash().Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	tx.Signature = signature
	start := time.Now()
	for i := 0; i < 100000; i++ {
		e, err := verifyTx(tx)
		if err != nil {
			log.Fatal(err)
		}
		if !e {
			log.Fatal("wrong signature")
		}
	}

	fmt.Println(100000 / time.Since(start).Seconds()) // true
}
