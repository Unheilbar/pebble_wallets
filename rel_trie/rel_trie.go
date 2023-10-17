package rel_trie

import "github.com/ethereum/go-ethereum/common"

// storageUpdated maps updates of storage to radix trie of slots, changed during update
type storageUpdated struct {
	storageChanges map[common.Hash]common.Hash
}

type currentUpdatesStorage map[common.Hash]common.Hash

// relTrie maps slotHashMaps to storageUpdates
type relTrie struct {
	slotHashMap map[common.Hash]storageUpdated
}

func (rel *relTrie) getState(key common.Hash) common.Hash {
	return rel.slotHashMap[key].storageChanges[key]
}

func (rel *relTrie) setState(storageRoot common.Hash, key common.Hash, val common.Hash) common.Hash {
	// return storage
	return common.Hash{}
}

// implements for evm StateDB
// GetState(common.Address, common.Hash) common.Hash
// SetState(common.Address, common.Hash, common.Hash)

func newRelTrie(prealloc int) *relTrie {
	rel := &relTrie{}

	rel.slotHashMap = make(map[common.Hash]storageUpdated, 100)
	rel.slotHashMap[common.Hash{}] = newStorageUpdated(2)

	return rel
}

func newStorageUpdated(prealloc int) storageUpdated {
	var strUpdated storageUpdated

	strUpdated.storageChanges = make(map[common.Hash]common.Hash, 2)
	return strUpdated
}
