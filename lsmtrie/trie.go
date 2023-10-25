package mlsmtrie

import (
	"fmt"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/Unheilbar/pebbke_wallets/binding"
	"github.com/Unheilbar/pebbke_wallets/core"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/Unheilbar/pebbke_wallets/core/vm"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
)

func New() {

}

var from = vm.AccountRef(common.HexToAddress("0x123"))
var to = common.HexToAddress("0x321")
var fakeWallet1 = "walletOne"
var fakeWallet2 = "walletTwo"

func runFakeVM() {
	state := NewFakeState()
	evm := vm.NewEVM(core.NewEVMBlockContext(&types.Header{Number: big.NewInt(1), Time: 123}), newTxContext(from.Address()), state, vm.DefaultCancunConfig())
	code := common.Hex2Bytes(binding.LsmBin[2:])

	_, accAddr, _, err := evm.Create(from, code, math.MaxBig256.Uint64(), new(big.Int))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(accAddr)
	emissions := getEmissions(25134)
	start := time.Now()
	interpreter := evm.Interpreter()
	contr := vm.NewContract(from, vm.AccountRef(accAddr), new(big.Int), math.MaxUint64)
	contr.SetCallCode(&accAddr, state.GetCodeHash(accAddr), state.GetCode(accAddr))
	for _, emission := range emissions {
		_, err := interpreter.Run(contr, emission, false)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("run", time.Since(start), len(state.state[accAddr].storage))
}

func getEmissions(size int) [][]byte {
	emissions := make([][]byte, 0)
	for i := 0; i < size; i++ {
		var arg [32]byte
		copy(arg[:], crypto.Keccak256Hash([]byte(fmt.Sprintf("%s%d", fakeWallet1, i))).Bytes()[:])
		input := packTX("emission", arg, big.NewInt(123))
		emissions = append(emissions, input)
	}

	return emissions
}

func packTX(method string, params ...interface{}) []byte {
	bind := binding.LsmABI

	pabi, err := abi.JSON(strings.NewReader(bind))

	if err != nil {
		log.Fatal("failed to read abi", err)
	}

	input, err := pabi.Pack(method, params...)
	if err != nil {
		log.Fatal("failed to pack tx", err)
	}

	return input
}

func newTxContext(from common.Address) vm.TxContext {
	return vm.TxContext{
		Origin:     from,
		GasPrice:   new(big.Int),
		BlobHashes: []common.Hash{},
	}
}
