package test_minter

import (
	"math/rand"
	"testing"
	"time"

	"github.com/Unheilbar/pebbke_wallets/tests/chad_v2"
	"github.com/ethereum/go-ethereum/common"
)

const (
	walletsAmount   = 10000
	transfersAmount = 100000

	rps = 100
)

func Test__Stess(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	tester := chad_v2.New(common.Address{}, nil)

	tester.InitAccs(walletsAmount, 0)

	tester.InitEmissions()

	tester.InitTransfers(transfersAmount)
}
