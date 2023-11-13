package eth

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"net"
	"time"

	"github.com/Unheilbar/pebble_wallets/core"
	"github.com/Unheilbar/pebble_wallets/core/types"
	pb "github.com/Unheilbar/pebble_wallets/proto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/rlp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Backend interface {
	SendTx(ctx context.Context, signedTx *types.Transaction) (time.Time, error)
	SubscribeChainHeadEvent(ch chan<- core.ChainHeadEvent) event.Subscription
	GetBalance([]byte, common.Address) *big.Int
}

type Server struct {
	ethApi Backend
	pb.UnimplementedPebbleAPIServer
}

func NewServer(e Backend, host string, port string) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	implServer := &Server{ethApi: e}
	pb.RegisterPebbleAPIServer(s, implServer)
	go s.Serve(listener)
	return nil
}

func (s *Server) SendTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionReply, error) {
	var tx = &types.Transaction{}

	err := rlp.DecodeBytes(in.Rlp, tx)
	if err != nil {
		log.Fatal(err)
	}

	tx.Hash()
	tx.WithSignature(in.Signature)

	recieveTime, err := s.ethApi.SendTx(ctx, tx)

	// PEBBLE Here we can implement error to status code later
	if err != nil {
		log.Fatal(err)
	}

	return &pb.TransactionReply{
		Status:       1,
		RecievedTime: uint64(recieveTime.UnixNano()),
	}, nil
}

func (s *Server) SubscribeBlocks(req *pb.SubscribeRequest, stream pb.PebbleAPI_SubscribeBlocksServer) error {
	ch := make(chan core.ChainHeadEvent)
	sub := s.ethApi.SubscribeChainHeadEvent(ch)
	defer sub.Unsubscribe()
	defer close(ch)
	for {
		select {
		case <-stream.Context().Done():
			log.Println("client closed stream connection")
			return status.Error(codes.Canceled, "stream has ended")
		case err := <-sub.Err():
			log.Println("client closed stream connection", err)
			return status.Error(codes.Canceled, "stream has ended")
		case event := <-ch:
			bblock, err := rlp.EncodeToBytes(event.Block)
			if err != nil {
				log.Println("block encode err", err)
				return status.Error(codes.Canceled, "stream has ended")
			}
			err = stream.Send(&pb.Block{
				BlockId:  event.Block.Number().Int64(),
				BlockRLP: bblock,
			})
			if err != nil {
				log.Println(err)
				return status.Error(codes.Canceled, "stream has ended")
			}
		}
	}
}

func (s *Server) GetBalance(ctx context.Context, req *pb.GetBalanceRequest) (*pb.BalanceReply, error) {
	addr := common.BytesToAddress(req.Address)
	res := s.ethApi.GetBalance(req.WalletId, addr)
	if res == nil {
		res = new(big.Int)
	}
	return &pb.BalanceReply{
		WalletId: req.WalletId,
		Balance:  res.Int64(),
	}, nil
}
