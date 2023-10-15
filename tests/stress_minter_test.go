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
	"github.com/ethereum/go-ethereum/crypto"
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

	runStress(ctx, api)
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

var minterStressWalletsAmount = 100000    // 100,000 wallets sounds good for a transaction
var minterStressTransfersAmount = 1000000 // million transactions

func runStress(ctx context.Context, api *eth.EthAPIBackend) {
	var tester chad.Chad
	tester.GenerateAccs(minterStressWalletsAmount)
	defaultDeployHash := crypto.Keccak256Hash([]byte("eventually will have to change it"))
	api.SendTx(context.Background(), tester.GetContractDeployTX(tester.GetTestAccByID(3).Address, common.Hex2Bytes(binding.StateMetaData.Bin[2:]), defaultDeployHash))
	time.Sleep(time.Second * 2) // wait to apply
	api.SendTx(context.Background(), tester.GetContractDeployTX(tester.GetTestAccByID(4).Address, common.Hex2Bytes(binding.EventsMetaData.Bin[2:]), defaultDeployHash))
	time.Sleep(time.Second * 2) // wait to apply
	api.SendTx(context.Background(), tester.GetContractDeployTX(tester.GetTestAccByID(5).Address, common.Hex2Bytes(binding.TransferMetaData.Bin[2:]), defaultDeployHash))
	time.Sleep(time.Second * 2) // wait to apply
	api.SendTx(context.Background(), tester.GetContractDeployTX(tester.GetTestAccByID(6).Address, common.Hex2Bytes(binding.ProxyMetaData.Bin[2:]), defaultDeployHash))
	time.Sleep(time.Second * 60) // wait for tx to apply
	var State = crypto.CreateAddress(tester.GetTestAccByID(3).Address, 0)
	var Transfer = crypto.CreateAddress(tester.GetTestAccByID(5).Address, 0)
	var Event = crypto.CreateAddress(tester.GetTestAccByID(4).Address, 0)
	var Proxy = crypto.CreateAddress(tester.GetTestAccByID(6).Address, 0)

	fmt.Printf("deployed addresses state %s transfer %s event %s proxy %s", State.Hex(), Transfer.Hex(), Event.Hex(), Proxy.Hex())

	// deployed addresses state 0x1aEa632C29D2978A5C6336A3B8BFE9d737EB8fE3 transfer 0x98aCaC3B9c77c934C12780a2852A959E674970A3 event 0x94a562Ef266F41D4AC4b125c1C2a5aAf7E952467 proxy 0x4BD6080baB7FB15D17bb211e333A87B7edE02D91
	emissions := tester.GenerateAccEmissionsTx(Proxy)
	api.SendTxs(context.Background(), emissions)
	time.Sleep(time.Second * 100) // wait emissions to process
	transfers := tester.GenerateTransfers(minterStressTransfersAmount, Proxy)
	api.SendTxs(context.Background(), transfers)
}
