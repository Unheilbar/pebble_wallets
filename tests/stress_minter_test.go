package tests

import (
	"context"
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
	api.SendTx(context.Background(), chad.GetContractDeployTX(common.Hex2Bytes(binding.StorageMetaData.Bin[2:])))

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

var minterStressWalletsAmount = 10000
var contrAddr = common.HexToAddress("0x5a443704dd4B594B382c22a083e2BD3090A6feF3")

func runStress(api *eth.EthAPIBackend) {
	var tester chad.Chad

	tester.GenerateAccs(minterStressWalletsAmount)
	emissions := tester.GenerateAccEmissionsTx(contrAddr)
	for _, tx := range emissions {
		api.SendTx(context.Background(), tx)
	}

	time.Sleep(time.Second * 10) // wait for tx to apply
}
