package main

import (
	"github.com/Unheilbar/pebbke_wallets/eth"
	"github.com/Unheilbar/pebbke_wallets/node"
)

func main() {
	makeFullNode()
}

func makeFullNode() (*node.Node, *eth.Ethereum) {
	ethService := eth.New()
	stack := node.New()

	return stack, ethService
}
