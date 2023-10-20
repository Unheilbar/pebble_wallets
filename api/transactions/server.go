package api

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/Unheilbar/pebbke_wallets/proto"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc"
)

type ethApi interface {
	SendTx(ctx context.Context, signedTx *types.Transaction)
}

type Server struct {
	ethApi ethApi
	proto.UnimplementedApiTransactionsServer
}

func NewServer(e ethApi, host string, port string) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	implServer := &Server{ethApi: e}
	proto.RegisterApiTransactionsServer(s, implServer)
	go s.Serve(listener)
	return nil
}

func (s *Server) SendTransaction(ctx context.Context, tx *proto.TransactionRequest) (*proto.TransactionReply, error) {
	to := common.BytesToAddress(tx.To)
	signedTx := types.NewTx(types.TxData{
		From:      common.BytesToAddress(tx.From),
		To:        &to,
		Id:        common.BytesToHash(tx.Id),
		Signature: tx.Signature,
		Data:      tx.Data,
	})
	s.ethApi.SendTx(ctx, signedTx)

	return &proto.TransactionReply{
		Status:       1,
		RecievedTime: uint64(time.Now().Unix()),
	}, nil
}
