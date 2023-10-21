package eth

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	pb "github.com/Unheilbar/pebbke_wallets/proto"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ethApi interface {
	SendTx(ctx context.Context, signedTx *types.Transaction) (time.Time, error)
}

type Server struct {
	ethApi ethApi
	pb.UnimplementedPebbleAPIServer
}

func NewServer(e ethApi, host string, port string) error {
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

func (s *Server) SendTransaction(ctx context.Context, tx *pb.TransactionRequest) (*pb.TransactionReply, error) {
	to := common.BytesToAddress(tx.To)
	signedTx := types.NewTx(types.TxData{
		From:      common.BytesToAddress(tx.From),
		To:        &to,
		Id:        common.BytesToHash(tx.Id),
		Signature: tx.Signature,
		Data:      tx.Data,
	})

	recieveTime, err := s.ethApi.SendTx(ctx, signedTx)

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
	for {
		select {
		case <-stream.Context().Done():
			log.Println("client closed stream connection")
			return status.Error(codes.Canceled, "stream has ended")
		default:
			err := stream.Send(&pb.Block{})
			if err != nil {
				log.Println(err)
				return status.Error(codes.Canceled, "stream has ended")
			}
		}
	}
}
