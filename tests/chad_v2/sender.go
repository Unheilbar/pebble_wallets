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

type timings struct {
	mu  sync.Mutex
	res map[common.Hash]*result

	maxSla            time.Duration
	triggerSla        time.Duration
	txAboveTriggerSla atomic.Int64
}

func (t *timings) setSend(id common.Hash, sentTime time.Time) {
	t.mu.Lock()
	t.res[id] = &result{nodeRecieve: sentTime}
	t.mu.Unlock()
}

func (t *timings) setRecieve(id common.Hash, recieveTime time.Time) {
	t.mu.Lock()
	if res, ok := t.res[id]; ok {
		res.clientRecieve = recieveTime
		sla := res.clientRecieve.Sub(res.nodeRecieve)
		if sla > t.maxSla {
			t.maxSla = sla
		}
		if sla > t.triggerSla {
			fmt.Println("here")
			t.txAboveTriggerSla.Add(1)
		}
	}
	t.mu.Unlock()
}

type result struct {
	nodeRecieve   time.Time
	clientRecieve time.Time
}

type Sender struct {
	client     pb.PebbleAPIClient
	testerData *Chad
	mu         sync.Mutex

	arrivedTxes map[common.Hash]time.Time

	emissionsResults map[common.Hash]result

	//counters
	reqCounter     atomic.Int64
	recieveCounter atomic.Int64

	//result
	result timings

	accsAmount      int
	transfersAmount int
}

func NewSender(nodeURL string, chad *Chad, triggerSla time.Duration) *Sender {
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
		result: timings{
			res:        make(map[common.Hash]*result),
			triggerSla: triggerSla,
		},
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
	lim := rate.NewLimiter(r, 1)
	emChan := make(chan txWithSignature, threads)
	s.initEmptyResult()
	ctx, cancel := context.WithCancel(context.Background())
	go s.display(ctx, "emission", time.Millisecond*50)
	go s.emissionsQueue(emChan)
	for i := 0; i < threads; i++ {
		go s.sendEmissions(emChan, lim)
	}
	s.waitEmissionsDone()
	close(emChan)
	cancel()
}

func (s *Sender) RunTransfers(rps int, threads int) {
	d := int(time.Second) / rps
	r := rate.Every(time.Duration(d))
	lim := rate.NewLimiter(r, 1)
	trChan := make(chan txWithSignature, threads)
	s.initEmptyResult()
	ctx, cancel := context.WithCancel(context.Background())
	go s.display(ctx, "transfer", time.Millisecond*50)
	go s.transferQueue(trChan)
	for i := 0; i < threads; i++ {
		go s.sendTransfers(trChan, lim)
	}
	s.waitTransfersDone()
	close(trChan)
	cancel()
}

func (s *Sender) transferQueue(ch chan txWithSignature) {
	for _, tx := range s.testerData.transfers {
		ch <- txWithSignature{tx.transaction, tx.signature}
	}
}

func (s *Sender) emissionsQueue(ch chan txWithSignature) {
	for _, tx := range s.testerData.emissions {
		ch <- txWithSignature{tx.transaction, tx.signature}
	}
}

func (s *Sender) sendEmissions(ch chan txWithSignature, lim *rate.Limiter) {
	for tx := range ch {
		lim.Wait(context.Background())
		reply, err := s.client.SendTransaction(context.Background(), txToProto(tx.tx, tx.signature))
		s.result.setSend(tx.tx.Id(), time.Unix(0, int64(reply.RecievedTime)))
		s.reqCounter.Add(1)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *Sender) sendTransfers(ch chan txWithSignature, lim *rate.Limiter) {
	for tx := range ch {
		lim.Wait(context.Background())
		reply, err := s.client.SendTransaction(context.Background(), txToProto(tx.tx, tx.signature))
		if err != nil || reply.Status != 1 {
			log.Fatal("err send transfer", err, tx.tx.Id().Hex())
		}
		s.result.setSend(tx.tx.Id(), time.Unix(0, int64(reply.RecievedTime)))
		s.reqCounter.Add(1)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s *Sender) waitEmissionsDone() {
	for {
		time.Sleep(time.Millisecond * 50)
		if s.recieveCounter.Load() == int64(s.accsAmount) {
			break
		}
	}
}

func (s *Sender) waitTransfersDone() {
	for {
		time.Sleep(time.Millisecond * 50)
		if s.recieveCounter.Load() == int64(s.transfersAmount) {
			break
		}
	}
}

func (s *Sender) display(ctx context.Context, operation string, interval time.Duration) {
	start := time.Now()
	for {
		select {
		case <-ctx.Done():
			fmt.Println()
			return
		case <-time.After(interval):
			dur := time.Since(start)
			rps := float64(s.reqCounter.Load()) / dur.Seconds()
			tps := float64(s.recieveCounter.Load()) / dur.Seconds()
			fmt.Printf("\r %s rps:%.2f tps:%.2f max sla: %s tx sla error: %d", operation, rps, tps, s.result.maxSla, s.result.txAboveTriggerSla.Load())
		}
	}

}

func (s *Sender) initEmptyResult() {
	s.recieveCounter.Store(0)
	s.reqCounter.Store(0)

	s.result = timings{
		res:        make(map[common.Hash]*result),
		triggerSla: s.result.triggerSla,
	}
}

func (s *Sender) updateStats(txs []*types.Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()
	recieveTime := time.Now()
	for _, tx := range txs {
		if _, ok := s.testerData.emissionsId[tx.Id()]; ok {
			s.recieveCounter.Add(1)
		} else if _, ok := s.testerData.transfersId[tx.Id()]; ok {
			s.recieveCounter.Add(1)
		}
		s.result.setRecieve(tx.Id(), recieveTime)
		s.arrivedTxes[tx.Id()] = recieveTime
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
