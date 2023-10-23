package chad_v2

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	pb "github.com/Unheilbar/pebbke_wallets/proto"
	"github.com/ethereum/go-ethereum/common"
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

	//counters
	emissionReqCounter atomic.Int64
	transferReqCounter atomic.Int64

	emissionRecieveCounter atomic.Int64
	transferRecieveCounter atomic.Int64

	accsAmount      int
	transfersAmount int
}

func NewSender(nodeURL string, chad *Chad) *Sender {
	conn, err := grpc.Dial(nodeURL, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}...)
	if err != nil {
		log.Fatal(err)
	}

	pebbleClient := pb.NewPebbleAPIClient(conn)

	return &Sender{
		client:          pebbleClient,
		testerData:      chad,
		arrivedTxes:     make(map[common.Hash]time.Time),
		accsAmount:      len(chad.accList),
		transfersAmount: len(chad.transfers),
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

			s.updateStats(block.Transactions)
		}
	}
}

func (s *Sender) Deploy() {
	for i, tx := range s.testerData.deploys {
		fmt.Printf("\rdeploy %d out of %d ...", i+1, len(s.testerData.deploys))
		s.mustDeploy(tx)
	}
	fmt.Println()
}

func (s *Sender) mustDeploy(tx *types.Transaction) {
	s.client.SendTransaction(context.Background(), txToProto(tx, []byte{1, 2, 3})) // deploy doesnt check signature
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

type txWithSignature struct {
	tx        *types.Transaction
	signature []byte
}

func (s *Sender) RunEmissions(rps int, threads int) {
	d := int(time.Second) / rps
	r := rate.Every(time.Duration(d))
	lim := rate.NewLimiter(r, rps/2)
	emChan := make(chan txWithSignature, threads)
	go s.emissionsQueue(emChan)
	for i := 0; i < threads; i++ {
		go s.sendEmissions(emChan, lim)
	}
	s.waitEmissionsDone()
	close(emChan)
}

func (s *Sender) emissionsQueue(ch chan txWithSignature) {
	for _, tx := range s.testerData.emissions {
		ch <- txWithSignature{tx.transaction, tx.signature}
	}
}

func (s *Sender) sendEmissions(ch chan txWithSignature, lim *rate.Limiter) {
	for tx := range ch {
		lim.Wait(context.Background())
		_, err := s.client.SendTransaction(context.Background(), txToProto(tx.tx, tx.signature))
		s.emissionReqCounter.Add(1)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *Sender) waitEmissionsDone() {
	for {
		time.Sleep(time.Millisecond * 50)
		if s.emissionRecieveCounter.Load() == int64(s.accsAmount) {
			break
		}
	}
}

func (s *Sender) Display(interval time.Duration) {

}

func (s *Sender) updateStats(txs []*types.Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, tx := range txs {
		fmt.Println(i, tx.Id(), tx.Time().UnixMilli())
		s.arrivedTxes[tx.Id()] = tx.Time()
	}
}

func txToProto(tx *types.Transaction, signature []byte) *pb.TransactionRequest {
	rlp, err := rlp.EncodeToBytes(tx)
	if err != nil {
		log.Fatal(err)
	}

	return &pb.TransactionRequest{
		Rlp:       rlp,
		Signature: signature,
	}
}
