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

	// start peers
	bootstrapNodes := []string{
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5001",
	}

	// node 1
	firstNode, firstEth := makeFullNode()
	service1 := RegisterRaftService(firstNode, firstEth, 1, bootstrapNodes, "./firstRaftNode")
	service1.StartRaftNode()
	service1.StartMinter()

	// node 2
	secondNode, secondEth := makeFullNode()
	service2 := RegisterRaftService(secondNode, secondEth, 2, bootstrapNodes, "./secondRaftNode")
	service2.StartRaftNode()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	api := eth.NewApi(firstEth)

	runStress(api)
	<-ctx.Done()
}

func makeFullNode() (*node.Node, *eth.Ethereum) {
	ethService := eth.New()
	stack := node.New()

	return stack, ethService
}
func RegisterRaftService(n *node.Node, e *eth.Ethereum, raftId uint16, bootstrapNodes []string, raftLogDir string) *raft.RaftService {
	service := raft.NewRaft(e, blockThrottle, raftId, bootstrapNodes, raftLogDir)
	log.Print("raft service registered")
	return service
}

var minterStressWalletsAmount = 10
var minterStressTransfersAmount = 10

func runStress(api *eth.EthAPIBackend) {
	var tester chad.Chad
	tester.GenerateAccs(minterStressWalletsAmount)
	api.SendTx(context.Background(), tester.GetContractDeployTX(tester.GetTestAccByID(3).From, common.Hex2Bytes(binding.StateMetaData.Bin[2:])))
	time.Sleep(time.Second * 2) // wait to apply
	api.SendTx(context.Background(), tester.GetContractDeployTX(tester.GetTestAccByID(4).From, common.Hex2Bytes(binding.EventsMetaData.Bin[2:])))
	time.Sleep(time.Second * 2) // wait to apply
	api.SendTx(context.Background(), tester.GetContractDeployTX(tester.GetTestAccByID(5).From, common.Hex2Bytes(binding.TransferMetaData.Bin[2:])))
	time.Sleep(time.Second * 2) // wait to apply
	api.SendTx(context.Background(), tester.GetContractDeployTX(tester.GetTestAccByID(6).From, common.Hex2Bytes(binding.ProxyMetaData.Bin[2:])))
	time.Sleep(time.Second * 2) // wait for tx to apply
	var State = crypto.CreateAddress(tester.GetTestAccByID(3).From, 0)
	var Transfer = crypto.CreateAddress(tester.GetTestAccByID(5).From, 0)
	var Event = crypto.CreateAddress(tester.GetTestAccByID(4).From, 0)
	var Proxy = crypto.CreateAddress(tester.GetTestAccByID(6).From, 0)

	fmt.Printf("deployed addresses state %s transfer %s event %s proxy %s", State.Hex(), Transfer.Hex(), Event.Hex(), Proxy.Hex())

	// deployed addresses state 0x1aEa632C29D2978A5C6336A3B8BFE9d737EB8fE3 transfer 0x98aCaC3B9c77c934C12780a2852A959E674970A3 event 0x94a562Ef266F41D4AC4b125c1C2a5aAf7E952467 proxy 0x4BD6080baB7FB15D17bb211e333A87B7edE02D91
	emissions := tester.GenerateAccEmissionsTx(Proxy)

	api.SendTxs(context.Background(), emissions)
	time.Sleep(time.Second * 3) // wait emissions to process
	transfers := tester.GenerateTransfers(minterStressTransfersAmount, Proxy)
	fmt.Println(len(transfers))
	api.SendTxs(context.Background(), transfers)
}
