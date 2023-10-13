package raft

import (
	"time"

	"github.com/Unheilbar/pebbke_wallets/core"
	"github.com/Unheilbar/pebbke_wallets/core/txpool"
	"github.com/Unheilbar/pebbke_wallets/eth"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/event"
)

type RaftService struct {
	blockchain *core.Blockchain
	chainDb    ethdb.Database // Block chain database
	txPool     *txpool.TxPool
	// accountManager *accounts.Manager
	// downloader     *downloader.Downloader

	// raftProtocolManager *ProtocolManager
	// startPeers []*enode.Node

	// we need an event mux to instantiate the blockchain
	eventMux *event.TypeMux
	minter   *minter
	// nodeKey          *ecdsa.PrivateKey
	// calcGasLimitFunc func(block *types.Block) uint64
}

func NewRaft(e *eth.Ethereum, blockTime time.Duration) *RaftService {
	service := &RaftService{
		chainDb:    e.ChainDb(),
		blockchain: e.Blockchain(),
		txPool:     e.TxPool(),
	}

	service.minter = newMinter(service, blockTime)

	return service
}

func (service *RaftService) TxPool() *txpool.TxPool {
	return service.TxPool()
}

func (service *RaftService) Blockchain() *core.Blockchain {
	return service.blockchain
}

func (service *RaftService) ChainDb() ethdb.Database { return service.chainDb }
