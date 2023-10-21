package chad_v2

import (
	"crypto/ecdsa"
	"log"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"

	"github.com/Unheilbar/pebbke_wallets/binding"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

const genesisPrivate = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
const emissionTokens = 1000000000 // amount of tokens for emission
const transferTokens = 10         // amount of tokens for transfer
const methodEmission = "emission"
const methodTransfer = "transfer"

type NodeClient interface {
	SendTransaction()
}

type chadAcc struct {
	Address common.Address
	Private *ecdsa.PrivateKey
}

type fixtureAcc struct {
	sync.Mutex
	Address         common.Address
	Private         *ecdsa.PrivateKey
	WalletId        [32]byte
	ExpectedBalance *big.Int //expexcted balance gets calculated every time we send a new transfer
}

type fixtureTransfer struct {
	fromAccount *fixtureAcc
	toAccount   *fixtureAcc
	sendTime    time.Time
	transaction *types.Transaction
}

type fixtureEmission struct {
	emissionWallet [32]byte
	transaction    *types.Transaction
}

type Chad struct {
	proxyAddress common.Address
	accsMap      map[common.Address]*fixtureAcc
	accList      []*fixtureAcc

	transfers []*fixtureTransfer
	emissions []*fixtureEmission

	accounts map[common.Address]chadAcc

	client      NodeClient
	mu          sync.Mutex
	lastGenesis string

	deploys []*types.Transaction
}

// genesisOffset determines from which point we need generate contract addresses
func New(proxyAddress common.Address, client NodeClient) *Chad {
	c := &Chad{
		proxyAddress: proxyAddress,

		transfers: make([]*fixtureTransfer, 0),
		emissions: make([]*fixtureEmission, 0),
		accsMap:   make(map[common.Address]*fixtureAcc),
		accList:   make([]*fixtureAcc, 0),
		client:    client,
	}

	c.initDeploys()

	return c
}

func (c *Chad) InitAccs(walletsAmount int, genesisOffset int) {
	var firstAddress = genesisPrivate
	for i := 0; i < genesisOffset; i++ {
		firstAddress = string(crypto.Keccak256Hash([]byte(firstAddress)).Hex()[2:])
	}

	// now we have new first address, that we use for generating the rest of accounts
	for i := 0; i < walletsAmount; i++ {
		privateKey, err := crypto.HexToECDSA(firstAddress)
		if err != nil {
			log.Fatal("err gen private", err)
		}

		generatedAddr := common.BytesToAddress(crypto.Keccak256(crypto.FromECDSAPub(&privateKey.PublicKey)[1:])[12:])

		acc := fixtureAccFromAddr(generatedAddr, privateKey)

		// add accs to chad map and list
		c.accsMap[acc.Address] = acc
		c.accList = append(c.accList, acc)

		firstAddress = string(crypto.Keccak256Hash([]byte(firstAddress)).Hex()[2:]) // generate new private key
	}
}

func (c *Chad) InitEmissions() {
	for _, acc := range c.accList {
		payload := packTX(methodEmission, acc.Address, acc.WalletId, big.NewInt(emissionTokens))
		c.emissions = append(c.emissions, &fixtureEmission{
			emissionWallet: acc.WalletId,
			transaction:    c.getTx(acc, payload, c.proxyAddress),
		})
	}
}

func (c *Chad) InitTransfers(transfersAmount int) {
	for i := 0; i < transfersAmount; i++ {
		acc1, acc2 := c.getRandomAccsPair()
		payload := packTX(methodTransfer, acc1.WalletId, acc2.WalletId, big.NewInt(transferTokens))
		c.transfers = append(c.transfers, &fixtureTransfer{
			fromAccount: acc1,
			toAccount:   acc2,
			transaction: c.getTx(acc1, payload, c.proxyAddress),
		})
	}
}

func fixtureAccFromAddr(addr common.Address, key *ecdsa.PrivateKey) *fixtureAcc {
	var walletID [32]byte
	copy(walletID[:], addr.Hash().Bytes()[:32])
	var acc = &fixtureAcc{
		Address:         addr,
		Private:         key,
		WalletId:        walletID,
		ExpectedBalance: new(big.Int),
	}
	return acc
}

func packTX(method string, params ...interface{}) []byte {
	bind := binding.ProxyABI

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

func (c *Chad) getTx(acc *fixtureAcc, payload []byte, to common.Address) *types.Transaction {
	id := []byte(uuid.New().String())
	genTx := types.NewTx(types.TxData{
		From: acc.Address,
		To:   &to,
		Id:   crypto.Keccak256Hash(id),
		Data: payload,
	})
	sign, err := crypto.Sign(genTx.Hash().Bytes(), acc.Private)
	if err != nil {
		log.Fatal("sign err", err)
	}
	return genTx.WithSignature(sign)
}

func (c *Chad) getRandomAccsPair() (*fixtureAcc, *fixtureAcc) {
	first := rand.Intn(len(c.accList))
	second := first
	for second == first {
		second = rand.Intn(len(c.accList))
	}
	return c.accList[first], c.accList[second]
}

func (c *Chad) initDeploys() {
	stateDeploy := newDeploy(c.accList[3].Address, common.Hex2Bytes(binding.StateMetaData.Bin[2:]))
	eventsDeploy := newDeploy(c.accList[4].Address, common.Hex2Bytes(binding.EventsMetaData.Bin[2:]))
	transferDeploy := newDeploy(c.accList[5].Address, common.Hex2Bytes(binding.TransferMetaData.Bin[2:]))
	proxyDeploy := newDeploy(c.accList[6].Address, common.Hex2Bytes(binding.ProxyMetaData.Bin[2:]))

	c.deploys = []*types.Transaction{stateDeploy, eventsDeploy, transferDeploy, proxyDeploy}
}

var defaultDeployHash = crypto.Keccak256Hash([]byte("eventually will have to change it"))

func newDeploy(deployer common.Address, code []byte) *types.Transaction {
	genTx := types.NewTx(types.TxData{
		From: deployer,
		To:   nil,
		Id:   defaultDeployHash,
		Data: code,
	})
	return genTx
}
