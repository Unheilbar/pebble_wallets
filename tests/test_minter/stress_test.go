package test_minter

import (
	"context"
	"math/rand"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/Unheilbar/pebbke_wallets/tests/chad_v2"
	"github.com/ethereum/go-ethereum/common"
)

const (
	walletsAmount   = 10
	transfersAmount = 10

	rps     = 5
	threads = 10
)

func Test__Stess(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testerData := chad_v2.New(common.HexToAddress("0x94a562Ef266F41D4AC4b125c1C2a5aAf7E952467"))

	testerData.InitAccs(walletsAmount, 0)
	testerData.InitDeploys()
	testerData.InitEmissions()
	testerData.InitTransfers(transfersAmount)

	sender := chad_v2.NewSender("localhost:6050", testerData)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go sender.Listen(ctx)

	sender.Deploy()
	sender.RunEmissions(rps, threads)
	<-ctx.Done()
}
