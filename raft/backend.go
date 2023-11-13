package raft

import (
	"time"

	"github.com/Unheilbar/pebble_wallets/core"
	"github.com/Unheilbar/pebble_wallets/core/txpool"
	"github.com/Unheilbar/pebble_wallets/core/types"
	"github.com/Unheilbar/pebble_wallets/eth"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
)

type RaftService struct {
	blockchain *core.Blockchain
	chainDb    ethdb.Database // Block chain database
	txPool     *txpool.TxPool

	eventMux *event.TypeMux
	minter   *minter
	raftNode *raftNode
	// nodeKey          *ecdsa.PrivateKey
	// calcGasLimitFunc func(block *types.Block) uint64
}

func NewRaft(e *eth.Ethereum, blockTime time.Duration, raftId uint16, bootstrapNodes []string, raftLogDir string) *RaftService {
	service := &RaftService{
		chainDb:    e.ChainDb(),
		blockchain: e.Blockchain(),
		txPool:     e.TxPool(),
	}

	proposeBlockC := make(chan *types.Block, 10)
	service.minter = newMinter(service, blockTime, proposeBlockC)

	service.raftNode = NewRaftNode(raftId, e.Blockchain(), proposeBlockC, bootstrapNodes, raftLogDir)
	return service
}

func (service *RaftService) TxPool() *txpool.TxPool {
	return service.txPool
}

func (service *RaftService) Blockchain() *core.Blockchain {
	return service.blockchain
}

func (service *RaftService) ChainDb() ethdb.Database { return service.chainDb }

func (service *RaftService) StartMinter() {
	service.minter.start()
}

func (service *RaftService) StartRaftNode() {
	service.raftNode.startRaft()
}
