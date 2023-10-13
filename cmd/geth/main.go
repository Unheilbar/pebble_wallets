package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"log"

	"github.com/Unheilbar/pebbke_wallets/eth"
	"github.com/Unheilbar/pebbke_wallets/node"
	"github.com/Unheilbar/pebbke_wallets/raft"
)

var blockThrottle = time.Millisecond * 50

func main() {
	f, err := os.OpenFile(fmt.Sprintf("./logs/node.log", time.Now().Second()), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Print("start")
	node, e := makeFullNode()
	RegisterRaftService(node, e)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
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
