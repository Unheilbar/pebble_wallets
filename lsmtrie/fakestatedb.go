package mlsmtrie

import (
	"fmt"

	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

type acc struct {
	code     []byte
	codeHash common.Hash
	storage  map[common.Hash]common.Hash
}

type fakestate struct {
	state map[common.Address]*acc
}

func NewFakeState() *fakestate {
	return &fakestate{
		state: make(map[common.Address]*acc, 10),
	}
}

func (f *fakestate) CreateAccount(a common.Address) {
	f.state[a] = &acc{
		storage: make(map[common.Hash]common.Hash, 25000),
	}
	fmt.Println("fake CreateAccount")
}

func (f *fakestate) GetCodeHash(c common.Address) common.Hash {
	//	fmt.Println("fake GetCodeHash")
	return f.state[c].codeHash
}

func (f *fakestate) GetCode(c common.Address) []byte {
	acc, ok := f.state[c]
	if !ok {
		return nil
	}

	return acc.code
}

func (f *fakestate) SetCode(c common.Address, code []byte) {
	acc := f.state[c]
	acc.code = code
}

func (f *fakestate) GetCodeSize(common.Address) int {
	//	fmt.Println("fake GetCodeSize")
	return 0
}

func (f *fakestate) AddRefund(uint64) {
	// fmt.Println("fake AddRefund")
}

func (f *fakestate) SubRefund(uint64) {
	// fmt.Println("fake SubRefund")
}

func (f *fakestate) GetRefund() uint64 {
	//fmt.Println("fake GetRefund")
	return 0
}

func (f *fakestate) GetCommittedState(common.Address, common.Hash) common.Hash {
	//fmt.Println("fake GetCommittedState")
	return common.Hash{}
}

func (f *fakestate) GetState(a common.Address, h common.Hash) common.Hash {
	//	fmt.Println("fake GetState")
	return f.state[a].storage[h]
}

func (f *fakestate) SetState(a common.Address, k common.Hash, v common.Hash) {
	//fmt.Println("fake SetState")
	f.state[a].storage[k] = v
}

func (f *fakestate) GetTransientState(addr common.Address, key common.Hash) common.Hash {
	//fmt.Println("fake GetTransientState")
	return common.Hash{}
}

func (f *fakestate) SetTransientState(addr common.Address, key, value common.Hash) {
	///fmt.Println("fake SetTransientState")
}

func (f *fakestate) SelfDestruct(common.Address) {
	//fmt.Println("fake SelfDestruct")
}
func (f *fakestate) HasSelfDestructed(common.Address) bool {
	//	fmt.Println("fake HasSelfDestructed")
	return false
}

func (f *fakestate) Selfdestruct6780(common.Address) {
	//fmt.Println("fake Selfdestruct6780")
}

// Exist reports whether the given account exists in state.
// Notably this should also return true for self-destructed accounts.
func (f *fakestate) Exist(addr common.Address) bool {
	acc, ok := f.state[addr]
	if !ok {
		return false
	}
	return len(acc.code) != 0
}

// Empty returns whether the given account is empty. Empty
// is defined according to EIP161 (balance = nonce = code = 0).
func (f *fakestate) Empty(common.Address) bool {
	//fmt.Println("fake Empty")
	return false
}

func (f *fakestate) AddressInAccessList(addr common.Address) bool {
	//fmt.Println("fake AddressInAccessList")
	return false
}
func (f *fakestate) SlotInAccessList(addr common.Address, slot common.Hash) (addressOk bool, slotOk bool) {
	//fmt.Println("fake SlotInAccessList")
	return false, false
}

// AddAddressToAccessList adds the given address to the access list. This operation is safe to perform
// even if the feature/fork is not active yet
func (f *fakestate) AddAddressToAccessList(addr common.Address) {
	//fmt.Println("fake AddAddressToAccessList")
}

// AddSlotToAccessList adds the given (address,slot) to the access list. This operation is safe to perform
// even if the feature/fork is not active yet
func (f *fakestate) AddSlotToAccessList(addr common.Address, slot common.Hash) {
	//fmt.Println("fake AddSlotToAccessList")
}
func (f *fakestate) Prepare(rules params.Rules, sender, coinbase common.Address, dest *common.Address, precompiles []common.Address, txAccesses types.AccessList) {
	//fmt.Println("fake Prepare")
}

func (f *fakestate) RevertToSnapshot(int) {
	//fmt.Println("fake RevertToSnapshot")
}
func (f *fakestate) Snapshot() int {
	//fmt.Println("fake Snapshot")
	return 0
}

func (f *fakestate) AddLog(*types.Log) {
	//fmt.Println("fake AddLog")
}
func (f *fakestate) AddPreimage(common.Hash, []byte) {
	//fmt.Println("fake AddPreimage")

}
