package test_minter

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"testing"
	"time"

	pb "github.com/Unheilbar/pebble_wallets/proto"
	"github.com/Unheilbar/pebble_wallets/tests/chad_v2"
	"github.com/Unheilbar/pebble_wallets/tests/client"
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
const clientUri = "192.168.0.1:6050"

var triggerSla = time.Millisecond * 1500

func Test__Stess(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	testerData := chad_v2.New(common.HexToAddress(defaultProxyAddress))

	testerData.InitAccs(walletsAmount, 0)
	testerData.InitDeploys()
	testerData.InitEmissions()
	testerData.InitTransfers(transfersAmount)

	sender := chad_v2.NewSender(clientUri, testerData, triggerSla)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go sender.Listen(ctx)

	sender.Deploy()
	sender.RunEmissions(rps, threads)
	sender.RunTransfers(rps, threads)
	time.Sleep(time.Second)
	hosts := []string{clientUri}
	checkAccounts := []int{10, 20, 30, 40, 50}
	for _, h := range hosts {
		client := client.New(clientUri)
		fmt.Println("check results host", h)
		for _, accID := range checkAccounts {
			acc := testerData.AccByID(accID)
			reply, err := client.GetBalance(context.Background(), &pb.GetBalanceRequest{
				WalletId: acc.WalletId[:],
				Address:  common.HexToAddress(defaultProxyAddress).Bytes(),
			})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("wallet ID", common.BytesToAddress(acc.WalletId[:]), "expected balance", acc.ExpectedBalance, "actualBalance", reply.Balance)
		}

	}

	<-ctx.Done()
}
