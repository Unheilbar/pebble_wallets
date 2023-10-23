package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"time"

	"log"

	"github.com/Unheilbar/pebbke_wallets/eth"
	"github.com/Unheilbar/pebbke_wallets/node"
	"github.com/Unheilbar/pebbke_wallets/raft"
)

var blockThrottle = time.Millisecond * 50

func main() {
	raftId := flag.Int("raftId", 0, "raft node id")
	bootstrapNodes := flag.String("bootstrapNodes", "http://127.0.0.1:5000,http://127.0.0.1:5001", "bootstrap nodes list")
	raftLog := flag.String("raftlog", "./secondRaftNode", "raft log path")
	chaindbPath := flag.String("chaindb", "./chaindb_2", "chain db path")
	logPath := flag.String("log", "./logs/node.log", "log save path")
	grpcHost := flag.String("grpc_host", "localhost", "host of grpc api")
	grpcPort := flag.String("grpc_port", "6050", "port of grpc api")
	miner := flag.Bool("miner", false, "miner")
	flag.Parse()

	f, err := os.OpenFile(*logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	node, e := makeFullNode(*chaindbPath, *grpcHost, *grpcPort)
	service := RegisterRaftService(node, e, uint16(*raftId), strings.Split(*bootstrapNodes, ","), *raftLog)
	service.StartRaftNode()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if *miner {
		service.StartMinter()
		fmt.Println("ya minter")
	}

	<-ctx.Done()
}

func makeFullNode(chaindbPath, grpcHost, grpcPort string) (*node.Node, *eth.Ethereum) {
	ethService := eth.New(chaindbPath, grpcHost, grpcPort)
	stack := node.New()

	return stack, ethService
}

func RegisterRaftService(n *node.Node, e *eth.Ethereum, raftId uint16, bootstrapNodes []string, raftLogDir string) *raft.RaftService {
	service := raft.NewRaft(e, blockThrottle, raftId, bootstrapNodes, raftLogDir)
	log.Println("raft service registered")
	return service
}
