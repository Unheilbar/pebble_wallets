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
	walletsAmount   = 100
	transfersAmount = 10

	rps     = 5
	threads = 10
)

func Test__Stess(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testerData := chad_v2.New(common.Address{})

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
