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
	"github.com/Unheilbar/pebbke_wallets/core"
	"github.com/Unheilbar/pebbke_wallets/eth"
	"github.com/Unheilbar/pebbke_wallets/node"
	"github.com/Unheilbar/pebbke_wallets/raft"
	"github.com/Unheilbar/pebbke_wallets/tests/chad"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	grpcApiHost = "127.0.0.1"
	grpcApiPort = "12345"
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
		// "http://127.0.0.1:5001",
	}

	// node 1
	firstNode, firstEth := makeFullNode("../chaindb_1")
	service1 := RegisterRaftService(firstNode, firstEth, 1, bootstrapNodes, "../firstRaftNode")
	service1.StartRaftNode()

	// node 2
	//secondNode, secondEth := makeFullNode("../chaindb_2")
	//service2 := RegisterRaftService(secondNode, secondEth, 2, bootstrapNodes, "../secondRaftNode")
	//service2.StartRaftNode()

	time.Sleep(5 * time.Second)
	service1.StartMinter()
	fmt.Println("started miner")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	api := eth.NewApi(firstEth)

	chainEvent := make(chan core.ChainHeadEvent)
	sub := firstEth.Blockchain().SubscribeChainHeadEvent(chainEvent)

	go func() {
		for {
			select {
			case event := <-chainEvent:
				fmt.Println("receieved block from chain event", event.Block.Number())
			case err := <-sub.Err():
				fmt.Println("err recieved from chain event", err)
				return
			}

		}
	}()
	runStress(ctx, api)
	<-ctx.Done()
}

func makeFullNode(dbPath string) (*node.Node, *eth.Ethereum) {
	ethService := eth.New(dbPath, grpcApiHost, grpcApiPort)
	stack := node.New()

	return stack, ethService
}
func RegisterRaftService(n *node.Node, e *eth.Ethereum, raftId uint16, bootstrapNodes []string, raftLogDir string) *raft.RaftService {
	service := raft.NewRaft(e, blockThrottle, raftId, bootstrapNodes, raftLogDir)
	log.Print("raft service registered")
	return service
}

var minterStressWalletsAmount = 10    // 1000 wallets sounds good for a transaction
var minterStressTransfersAmount = 100 // 10 million transactions

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
	time.Sleep(time.Second * 2) // wait for tx to apply
	var State = crypto.CreateAddress(tester.GetTestAccByID(3).Address, 1)
	var Transfer = crypto.CreateAddress(tester.GetTestAccByID(5).Address, 1)
	var Event = crypto.CreateAddress(tester.GetTestAccByID(4).Address, 1)
	var Proxy = crypto.CreateAddress(tester.GetTestAccByID(6).Address, 1)

	log.Printf("deployed addresses state %s transfer %s event %s proxy %s", State.Hex(), Transfer.Hex(), Event.Hex(), Proxy.Hex())
	fmt.Println()

	// deployed addresses state
	// 0x1aEa632C29D2978A5C6336A3B8BFE9d737EB8fE3 transfer

	// 0x98aCaC3B9c77c934C12780a2852A959E674970A3 event
	// 0x94a562Ef266F41D4AC4b125c1C2a5aAf7E952467 proxy
	// 0x4BD6080baB7FB15D17bb211e333A87B7edE02D91
	emissions := tester.GenerateAccEmissionsTx(Proxy)
	transfers := tester.GenerateTransfers(minterStressTransfersAmount, Proxy)
	time.Sleep(time.Second * 10) // wait minter to prepare
	for _, e := range emissions {
		api.SendTxs(context.Background(), e)
	}

	time.Sleep(time.Second * 10) // wait emissions to process

	fmt.Println("Generated transaction count", len(transfers))
	for _, t := range transfers {
		api.SendTxs(context.Background(), t)
	}
}
