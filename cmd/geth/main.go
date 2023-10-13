package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"log"

	"github.com/Unheilbar/pebbke_wallets/binding"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/Unheilbar/pebbke_wallets/eth"
	"github.com/Unheilbar/pebbke_wallets/node"
	"github.com/Unheilbar/pebbke_wallets/raft"
	"github.com/ethereum/go-ethereum/common"
)

var blockThrottle = time.Millisecond * 50

func main() {
	f, err := os.OpenFile("./logs/node.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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
	api.SendTx(context.Background(), getContractDeployTX(common.Hex2Bytes(binding.StorageMetaData.Bin[2:])))
	go runLoad(api)
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

// remove later
func getContractDeployTX(contrCode []byte) *types.Transaction {
	return &types.Transaction{
		From:  common.Address{},
		To:    common.Address{},
		Input: contrCode,
	}
}

func runLoad(api *eth.EthAPIBackend)
