package chad_v2

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	pb "github.com/Unheilbar/pebbke_wallets/proto"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/time/rate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type result struct {
	sent     time.Time
	recieved time.Time
}

type Sender struct {
	client     pb.PebbleAPIClient
	testerData *Chad
	mu         sync.Mutex

	arrivedTxes map[common.Hash]time.Time

	emissionsResults map[common.Hash]result
}

func NewSender(nodeURL string, chad *Chad) *Sender {
	conn, err := grpc.Dial(nodeURL, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}...)
	if err != nil {
		log.Fatal(err)
	}

	pebbleClient := pb.NewPebbleAPIClient(conn)

	return &Sender{
		client:      pebbleClient,
		testerData:  chad,
		arrivedTxes: make(map[common.Hash]time.Time),
	}
}

// Listen gets new blocks and updates actual states based on blocks events
func (s *Sender) Listen(ctx context.Context) error {
	stream, err := s.client.SubscribeBlocks(ctx, &pb.SubscribeRequest{})
	if err != nil {
		return err
	}
	defer stream.CloseSend()

	for {
		select {
		case <-ctx.Done():
			log.Println("context canceled")
			stream.CloseSend()
		default:
			msg, err := stream.Recv()
			if err != nil {
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

			s.mu.Lock()
			for _, tx := range block.Transactions {
				s.arrivedTxes[tx.Id()] = time.Now()
			}
			s.mu.Unlock()
		}
	}
}

func (s *Sender) Deploy() {
	for i, tx := range s.testerData.deploys {
		fmt.Printf("\rdeploy %d out of %d ...", i+1, len(s.testerData.deploys))
		s.mustDeploy(tx)
	}
}

func (s *Sender) mustDeploy(tx *types.Transaction) {
	s.client.SendTransaction(context.Background(), txToProto(tx))
	for {
		s.mu.Lock()
		_, ok := s.arrivedTxes[tx.Id()]
		if ok {
			delete(s.arrivedTxes, tx.Id())
			s.mu.Unlock()
			return
		}
		s.mu.Unlock()
		time.Sleep(time.Millisecond * 50)
	}
}

func (s *Sender) RunEmissions(rps int, threads int) {
	d := int(time.Second) / rps
	r := rate.Every(time.Duration(d))
	lim := rate.NewLimiter(r, rps/2)
	emChan := make(chan *types.Transaction, threads)
	go s.emissionsQueue(emChan)
	go s.sendEmissions(emChan, lim)
}

func (s *Sender) emissionsQueue(ch chan *types.Transaction) {
	for _, tx := range s.testerData.emissions {
		ch <- tx.transaction
	}
}

func (s *Sender) sendEmissions(ch chan *types.Transaction, lim *rate.Limiter) {
	for tx := range ch {
		lim.Wait(context.Background())
		_, err := s.client.SendTransaction(context.Background(), txToProto(tx))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func txToProto(tx *types.Transaction) *pb.TransactionRequest {
	rlp, err := rlp.EncodeToBytes(tx)
	if err != nil {
		log.Fatal(err)
	}
	if tx.To() != nil {
		fmt.Println(tx.Hash(), tx.From().Hex(), tx.To().Hex(), tx.Id().Hex(), crypto.Keccak256Hash(tx.Data()), crypto.Keccak256Hash(tx.Signature()))
	}

	return &pb.TransactionRequest{
		Rlp: rlp,
	}
}
