package miner

import (
	"sync"

	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/event"
)

type Miner struct {
	mux     *event.TypeMux
	engine  consensus.Engine
	exitCh  chan struct{}
	startCh chan struct{}
	stopCh  chan struct{}
	worker  *worker

	wg sync.WaitGroup
}
