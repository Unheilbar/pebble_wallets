package raft

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set"

	"github.com/Unheilbar/pebbke_wallets/core/types"

	"github.com/Unheilbar/pebbke_wallets/core"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/opt"

	"go.etcd.io/etcd/client/pkg/v3/fileutil"
	rafttypes "go.etcd.io/etcd/client/pkg/v3/types"
	etcdRaft "go.etcd.io/etcd/raft/v3"
	"go.etcd.io/etcd/raft/v3/raftpb"
	"go.etcd.io/etcd/server/v3/etcdserver/api/rafthttp"
	"go.etcd.io/etcd/server/v3/etcdserver/api/snap"
	stats "go.etcd.io/etcd/server/v3/etcdserver/api/v2stats"
	"go.etcd.io/etcd/server/v3/wal"
	"go.etcd.io/etcd/server/v3/wal/walpb"

	"go.uber.org/zap"
)

type raftNode struct {
	mu sync.RWMutex

	id             int
	join           bool
	bootstrapNodes []string

	// Blockchain services
	blockchain *core.Blockchain

	confState     raftpb.ConfState
	snapshotIndex uint64
	appliedIndex  uint64

	unsafeRawNode etcdRaft.Node
	raftStorage   *etcdRaft.MemoryStorage
	wal           *wal.WAL

	//raft snapshot
	snapdir          string
	snapshotter      *snap.Snapshotter
	snapshotterReady chan *snap.Snapshotter // signals when snapshotter is ready

	transport *rafthttp.Transport

	//raft log
	waldir string

	// storage
	quorumRaftdb *leveldb.DB

	blockProposalC chan *types.Block

	removedPeers mapset.Set // *Permanently removed* peers

	httpstopc chan struct{} // signals http server to shutdown
	httpdonec chan struct{} // signals http server shutdown complete
	quitSync  chan struct{}
	logger    *zap.Logger
}

func openQuorumRaftDb(path string) (db *leveldb.DB, err error) {
	// Open the db and recover any potential corruptions
	db, err = leveldb.OpenFile(path, &opt.Options{
		OpenFilesCacheCapacity: -1, // -1 means 0??
		BlockCacheCapacity:     -1,
	})
	if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
		db, err = leveldb.RecoverFile(path, nil)
	}
	return
}

func NewRaftNode(raftId uint16, blockchain *core.Blockchain, blockProposalC chan *types.Block, bootstrapNodes []string, raftLogDir string) *raftNode {
	waldir := fmt.Sprintf("%s/raft-wal", raftLogDir)
	snapdir := fmt.Sprintf("%s/raft-snap", raftLogDir)
	//init storage
	quorumRaftdb := fmt.Sprintf("%s/quorum-raft-state", raftLogDir)
	db, err := openQuorumRaftDb(quorumRaftdb)
	if err != nil {
		log.Fatal("failed open raft db err ", err)
	}

	return &raftNode{
		id:               int(raftId),
		blockchain:       blockchain,
		join:             false,
		bootstrapNodes:   bootstrapNodes,
		removedPeers:     mapset.NewSet(),
		snapdir:          snapdir,
		waldir:           waldir,
		quorumRaftdb:     db,
		raftStorage:      etcdRaft.NewMemoryStorage(),
		blockProposalC:   blockProposalC,
		logger:           zap.NewExample(),
		snapshotterReady: make(chan *snap.Snapshotter, 1),
	}
}

func (pm *raftNode) Start() {
	fmt.Println("starting raft manager")

	//todo add start manager
}

func (rn *raftNode) startRaft() {
	if !fileutil.Exist(rn.snapdir) {
		if err := os.Mkdir(rn.snapdir, 0750); err != nil {
			log.Fatalf("raftexample: cannot create dir for snapshot (%v)", err)
		}
	}

	rn.snapshotter = snap.New(zap.NewExample(), rn.snapdir)
	oldwal := wal.Exist(rn.waldir)
	rn.wal = rn.replayWAL()

	//reply finished
	rn.snapshotterReady <- rn.snapshotter

	rpeers := make([]etcdRaft.Peer, len(rn.bootstrapNodes))
	for i := range rpeers {

		rpeers[i] = etcdRaft.Peer{ID: uint64(i + 1)}
	}

	c := &etcdRaft.Config{
		ID:                        uint64(rn.id),
		ElectionTick:              10,
		HeartbeatTick:             1,
		Storage:                   rn.raftStorage,
		MaxSizePerMsg:             1024 * 1024,
		MaxInflightMsgs:           256,
		MaxUncommittedEntriesSize: 1 << 30,
	}

	if oldwal || rn.join {
		rn.unsafeRawNode = etcdRaft.RestartNode(c)
	} else {
		rn.unsafeRawNode = etcdRaft.StartNode(c, rpeers)
	}

	rn.transport = &rafthttp.Transport{
		Logger:      rn.logger,
		ID:          rafttypes.ID(rn.id),
		ClusterID:   0x1000,
		Raft:        rn,
		ServerStats: stats.NewServerStats("", ""),
		LeaderStats: stats.NewLeaderStats(zap.NewExample(), strconv.Itoa(rn.id)),
		ErrorC:      make(chan error),
	}

	rn.transport.Start()
	for i := range rn.bootstrapNodes {
		if i+1 != rn.id {
			rn.transport.AddPeer(rafttypes.ID(i+1), []string{rn.bootstrapNodes[i]})
		}
	}

	go rn.serveRaft()
	go rn.eventLoop()
	go rn.serveLocalProposals()
}

func (rc *raftNode) serveRaft() {
	url, err := url.Parse(rc.bootstrapNodes[rc.id-1])
	if err != nil {
		log.Fatalf("raftexample: Failed parsing URL (%v)", err)
	}

	ln, err := newStoppableListener(url.Host, rc.httpstopc)
	if err != nil {
		log.Fatalf("raftexample: Failed to listen rafthttp (%v)", err)
	}

	err = (&http.Server{Handler: rc.transport.Handler()}).Serve(ln)
	select {
	case <-rc.httpstopc:
	default:
		log.Fatalf("raftexample: Failed to serve rafthttp (%v)", err)
	}
	close(rc.httpdonec)
}

func (pm *raftNode) Process(ctx context.Context, m raftpb.Message) error {
	return pm.rawNode().Step(ctx, m)
}

func (pm *raftNode) IsIDRemoved(id uint64) bool {
	return pm.isRaftIdRemoved(uint16(id))
}

func (pm *raftNode) ReportUnreachable(id uint64) {
	log.Println("peer is currently unreachable", "peer id", id)

	pm.rawNode().ReportUnreachable(id)
}

func (pm *raftNode) ReportSnapshot(id uint64, status etcdRaft.SnapshotStatus) {
	if status == etcdRaft.SnapshotFailure {
		log.Println("failed to send snapshot", "raft peer", id)
	} else if status == etcdRaft.SnapshotFinish {
		log.Println("finished sending snapshot", "raft peer", id)
	}

	pm.rawNode().ReportSnapshot(id, status)
}

// Serve two channels to handle new blocks and raft configuration changes originating locally.
func (n *raftNode) serveLocalProposals() {
	for {
		select {
		case block, ok := <-n.blockProposalC:
			log.Println("serve local block", block.Number().Int64())

			if !ok {
				log.Println("error: read from blockProposalC failed")
				return
			}

			size, r, err := rlp.EncodeToReader(block)
			if err != nil {
				panic(fmt.Sprintf("error: failed to send RLP-encoded block: %s", err.Error()))
			}

			var buffer = make([]byte, uint32(size))
			r.Read(buffer)

			// blocks until accepted by the raft state machine
			err = n.rawNode().Propose(context.Background(), buffer)
			if err != nil {
				log.Println("error propose block ", err)
			}
		case <-n.quitSync:
			return
		}
	}
}

const raftTickerInterval int = 100

func (n *raftNode) eventLoop() {
	ticker := time.NewTicker(time.Duration(raftTickerInterval) * time.Millisecond)
	defer ticker.Stop()
	defer n.wal.Close()

	for {
		select {
		case <-ticker.C:
			n.rawNode().Tick()

		case rd := <-n.rawNode().Ready():
			n.wal.Save(rd.HardState, rd.Entries)
			if !etcdRaft.IsEmptySnap(rd.Snapshot) {
				n.saveSnap(rd.Snapshot)
				n.raftStorage.ApplySnapshot(rd.Snapshot)
				n.publishSnapshot(rd.Snapshot)
			}
			n.raftStorage.Append(rd.Entries)
			n.transport.Send(rd.Messages)

			for _, entry := range n.entriesToApply(rd.CommittedEntries) {
				switch entry.Type {
				case raftpb.EntryNormal:
					if len(entry.Data) == 0 {
						break
					}

					var block types.Block
					err := rlp.DecodeBytes(entry.Data, &block)
					if err != nil {
						log.Println("error decoding block err", err)
					}

					err = n.blockchain.InsertChain(&block)
					if err != nil {
						log.Fatal("error insert to blockchain node ", n.id, err)
					}

					elapsed := time.Since(time.Unix(0, int64(block.Header().Time)))
					committedTxes := len(block.Transactions)
					log.Println("🔨  Insert chain block", "node", n.id, "number", block.Number(), "hash", fmt.Sprintf("%x", block.Hash().Bytes()[:4]), "elapsed", elapsed, "len(txs): ", committedTxes, getSpeed(committedTxes, elapsed), "tx/s ")
				case raftpb.EntryConfChange:
					var cc raftpb.ConfChange
					cc.Unmarshal(entry.Data)
					log.Println("new cc ", cc.NodeID, cc.Context, cc.ID)
					n.confState = *n.rawNode().ApplyConfChange(cc)
					switch cc.Type {
					case raftpb.ConfChangeAddNode:
						if len(cc.Context) > 0 {
							log.Println("add new node with context")
							n.transport.AddPeer(rafttypes.ID(cc.NodeID), []string{string(cc.Context)})
						}
					case raftpb.ConfChangeRemoveNode:
						log.Println("I've been removed from the cluster! Shutting down.")
						//TODO add remove peers
					}
				}

				// after commit, update appliedIndex
				n.advanceAppliedIndex(entry.Index)
			}

			n.rawNode().Advance()
		}

	}
}

func getSpeed(committedTxes int, elapsed time.Duration) float64 {
	return float64(committedTxes) / elapsed.Seconds()
}

func (n *raftNode) rawNode() etcdRaft.Node {
	for n.unsafeRawNode == nil {
		time.Sleep(100 * time.Millisecond)
	}

	return n.unsafeRawNode
}

// replayWAL replays WAL entries into the raft instance.
func (n *raftNode) replayWAL() *wal.WAL {
	log.Printf("replaying WAL of member %d", n.id)
	snapshot := n.loadSnapshot()
	w := n.openWAL(snapshot)
	_, st, ents, err := w.ReadAll()
	if err != nil {
		log.Fatalf("raftexample: failed to read WAL (%v)", err)
	}
	n.raftStorage = etcdRaft.NewMemoryStorage()
	if snapshot != nil {
		n.raftStorage.ApplySnapshot(*snapshot)
	}
	n.raftStorage.SetHardState(st)

	// append to storage so raft starts at the right place in log
	n.raftStorage.Append(ents)

	return w
}

// openWAL returns a WAL ready for reading.
func (rc *raftNode) openWAL(snapshot *raftpb.Snapshot) *wal.WAL {
	if !wal.Exist(rc.waldir) {
		if err := os.Mkdir(rc.waldir, 0750); err != nil {
			log.Fatalf("raftexample: cannot create dir for wal (%v)", err)
		}

		w, err := wal.Create(zap.NewExample(), rc.waldir, nil)
		if err != nil {
			log.Fatalf("raftexample: create wal error (%v)", err)
		}
		w.Close()
	}

	walsnap := walpb.Snapshot{}
	if snapshot != nil {
		walsnap.Index, walsnap.Term = snapshot.Metadata.Index, snapshot.Metadata.Term
	}
	log.Printf("loading WAL at term %d and index %d", walsnap.Term, walsnap.Index)
	w, err := wal.Open(zap.NewExample(), rc.waldir, walsnap)
	if err != nil {
		log.Fatalf("raftexample: error loading wal (%v)", err)
	}

	return w
}

func (n *raftNode) loadSnapshot() *raftpb.Snapshot {
	if wal.Exist(n.waldir) {
		walSnaps, err := wal.ValidSnapshotEntries(n.logger, n.waldir)
		if err != nil {
			log.Fatalf("raftexample: error listing snapshots (%v)", err)
		}
		snapshot, err := n.snapshotter.LoadNewestAvailable(walSnaps)
		if err != nil && err != snap.ErrNoSnapshot {
			log.Fatalf("raftexample: error loading snapshot (%v)", err)
		}
		return snapshot
	}
	return &raftpb.Snapshot{}
}

func (pm *raftNode) isRaftIdRemoved(id uint16) bool {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return pm.removedPeers.Contains(id)
}

func (n *raftNode) saveSnap(snap raftpb.Snapshot) error {
	walSnap := walpb.Snapshot{
		Index:     snap.Metadata.Index,
		Term:      snap.Metadata.Term,
		ConfState: &snap.Metadata.ConfState,
	}
	// save the snapshot file before writing the snapshot to the wal.
	// This makes it possible for the snapshot file to become orphaned, but prevents
	// a WAL snapshot entry from having no corresponding snapshot file.
	if err := n.snapshotter.SaveSnap(snap); err != nil {
		return err
	}
	if err := n.wal.SaveSnapshot(walSnap); err != nil {
		return err
	}
	return n.wal.ReleaseLockTo(snap.Metadata.Index)
}

func (rc *raftNode) publishSnapshot(snapshotToSave raftpb.Snapshot) {
	if etcdRaft.IsEmptySnap(snapshotToSave) {
		return
	}

	log.Printf("publishing snapshot at index %d", rc.snapshotIndex)
	defer log.Printf("finished publishing snapshot at index %d", rc.snapshotIndex)

	if snapshotToSave.Metadata.Index <= rc.appliedIndex {
		log.Fatalf("snapshot index [%d] should > progress.appliedIndex [%d]", snapshotToSave.Metadata.Index, rc.appliedIndex)
	}

	//TODO add sync with

	rc.confState = snapshotToSave.Metadata.ConfState
	rc.snapshotIndex = snapshotToSave.Metadata.Index
	rc.appliedIndex = snapshotToSave.Metadata.Index
}

func (rc *raftNode) entriesToApply(ents []raftpb.Entry) (nents []raftpb.Entry) {
	if len(ents) == 0 {
		return ents
	}
	firstIdx := ents[0].Index
	if firstIdx > rc.appliedIndex+1 {
		log.Fatalf("first index of committed entry[%d] should <= progress.appliedIndex[%d]+1", firstIdx, rc.appliedIndex)
	}
	if rc.appliedIndex-firstIdx+1 < uint64(len(ents)) {
		nents = ents[rc.appliedIndex-firstIdx+1:]
	}
	return nents
}

// Sets new appliedIndex in-memory, *and* writes this appliedIndex to LevelDB.
func (n *raftNode) advanceAppliedIndex(index uint64) {
	n.writeAppliedIndex(index)

	n.mu.Lock()
	n.appliedIndex = index
	n.mu.Unlock()
}

var (
	appliedDbKey = []byte("applied")
)

var (
	noFsync = &opt.WriteOptions{
		NoWriteMerge: false,
		Sync:         false,
	}
)

func (n *raftNode) writeAppliedIndex(index uint64) {
	log.Println("persisted the latest applied index", "index", index)
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	n.quorumRaftdb.Put(appliedDbKey, buf, noFsync)
}
