package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/Unheilbar/pebbke_wallets/binding"
	"github.com/Unheilbar/pebbke_wallets/eth"
	"github.com/Unheilbar/pebbke_wallets/node"
	"github.com/Unheilbar/pebbke_wallets/raft"
	"github.com/Unheilbar/pebbke_wallets/tests/chad"
	"github.com/ethereum/go-ethereum/common"
)

var blockThrottle = time.Millisecond * 50

func Test__RunStressMinter(t *testing.T) {
	f, err := os.OpenFile("../logs/node.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	node, e := makeFullNode()
	RegisterRaftService(node, e)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	api := eth.NewApi(e)

	runStress(api)
	<-ctx.Done()
}

func makeFullNode() (*node.Node, *eth.Ethereum) {
	ethService := eth.New()
	stack := node.New()

	return stack, ethService
}
func RegisterRaftService(n *node.Node, e *eth.Ethereum) {
	raft.NewRaft(e, blockThrottle)

	log.Print("raft service registered")
}

var minterStressWalletsAmount = 1000
var minterStressTransfersAmount = 10000
var Proxy = common.HexToAddress("0x5dBC355B93DD7A0C0D759Fd6a7859d2610219221")
var Transfer = common.HexToAddress("0x6C02e060D0E1CAD7c039A9aE3aBc29A40b3DFF1f")
var State = common.HexToAddress("0xE11f8d55a93bF877a091a3C54C071AAc5cC0b01D")
var Event = common.HexToAddress("0x6027946B05e7ab6Ef245093622AB18eaD5453877")

func runStress(api *eth.EthAPIBackend) {
	var tester chad.Chad
	tester.GenerateAccs(minterStressWalletsAmount)
	api.SendTx(context.Background(), tester.GetContractDeployTX(2, common.Hex2Bytes(binding.StateMetaData.Bin[2:])))
	time.Sleep(time.Second * 2) // wait to apply
	api.SendTx(context.Background(), tester.GetContractDeployTX(3, common.Hex2Bytes(binding.EventsMetaData.Bin[2:])))
	time.Sleep(time.Second * 2) // wait to apply
	api.SendTx(context.Background(), tester.GetContractDeployTX(4, common.Hex2Bytes(binding.TransferMetaData.Bin[2:])))
	time.Sleep(time.Second * 2) // wait to apply
	api.SendTx(context.Background(), tester.GetContractDeployTX(5, common.Hex2Bytes(binding.ProxyMetaData.Bin[2:])))

	time.Sleep(time.Second * 2) // wait for tx to apply
	emissions := tester.GenerateAccEmissionsTx(Proxy)
	api.SendTxs(context.Background(), emissions)

	transfers := tester.GenerateTransfers(minterStressTransfersAmount, Proxy)
	fmt.Println(len(transfers))
	api.SendTxs(context.Background(), transfers)

}
