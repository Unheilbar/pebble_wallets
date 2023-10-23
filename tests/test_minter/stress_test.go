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
	walletsAmount   = 10000
	transfersAmount = 100000

	rps     = 3000
	threads = 100
)

// state 0x6027946B05e7ab6Ef245093622AB18eaD5453877
// event 0x6C02e060D0E1CAD7c039A9aE3aBc29A40b3DFF1f
// transfer 0x5dBC355B93DD7A0C0D759Fd6a7859d2610219221
// proxy 0xb84A87293A2f2A9387DcE04145bB5d97942c1129

const defaultProxyAddress = "0xb84A87293A2f2A9387DcE04145bB5d97942c1129"

var triggerSla = time.Millisecond * 1500

func Test__Stess(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testerData := chad_v2.New(common.HexToAddress(defaultProxyAddress))

	testerData.InitAccs(walletsAmount, 0)
	testerData.InitDeploys()
	testerData.InitEmissions()
	testerData.InitTransfers(transfersAmount)

	sender := chad_v2.NewSender("localhost:6050", testerData, triggerSla)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go sender.Listen(ctx)

	sender.Deploy()
	sender.RunEmissions(rps, threads)
	sender.RunTransfers(rps, threads)
	<-ctx.Done()
}
