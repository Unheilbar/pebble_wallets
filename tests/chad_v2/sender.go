package chad_v2

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	pb "github.com/Unheilbar/pebbke_wallets/proto"
	"github.com/ethereum/go-ethereum/rlp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Sender struct {
	client     pb.PebbleAPIClient
	testerData *Chad
}

func NewSender(nodeURL string, chad *Chad) *Sender {
	conn, err := grpc.Dial(nodeURL, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}...)
	if err != nil {
		log.Fatal(err)
	}

	pebbleClient := pb.NewPebbleAPIClient(conn)

	return &Sender{pebbleClient, chad}
}

// Listen gets new blocks and updates actual states based on blocks events
func (s *Sender) Listen(ctx context.Context) error {
	stream, err := s.client.SubscribeBlocks(ctx, &pb.SubscribeRequest{})
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("context canceled")
			stream.CloseSend()
		default:
			msg, err := stream.Recv()
			if err != nil {
				stream.CloseSend()
				log.Println("recieve error from stream", err)
				return err
			}

			if msg == nil {
				log.Println("recieved nil from stream")
				return errors.New("nil msg from stream")
			}

			var block types.Block

			err = rlp.DecodeBytes(msg.BlockRLP, &block)
			if err != nil {
				log.Fatal("err decoding recieved block", err)
			}

			fmt.Println("recieved block", block.Header().Root)
		}
	}
}

func (s *Sender) Deploy() {
	for i, tx := range s.testerData.deploys {
		fmt.Printf("deploy %d out of %d ...", i+1, len(s.testerData.deploys))
		s.client.SendTransaction(context.Background(), txToProto(tx))
	}
}

func txToProto(tx *types.Transaction) *pb.TransactionRequest {
	var to []byte
	if tx.To() != nil {
		to = tx.To().Bytes()
	}
	return &pb.TransactionRequest{
		From:      tx.From().Bytes(),
		To:        to,
		Id:        tx.Id().Bytes(),
		Signature: tx.Signature(),
		Data:      tx.Data(),
	}
}
