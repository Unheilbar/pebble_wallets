package chad_v2

import (
	"bytes"
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
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/google/uuid"
)

const genesisPrivate = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
const emissionTokens = 1000000000 // amount of tokens for emission
const transferTokens = 10         // amount of tokens for transfer
const methodEmission = "emission"
const methodTransfer = "transfer"

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
	signature   []byte
}

type fixtureEmission struct {
	emissionWallet [32]byte
	transaction    *types.Transaction
	signature      []byte
}

type Chad struct {
	proxyAddress common.Address
	accsMap      map[common.Address]*fixtureAcc
	accList      []*fixtureAcc

	transfers []*fixtureTransfer
	emissions []*fixtureEmission

	emissionsId map[common.Hash]*fixtureEmission
	transfersId map[common.Hash]*fixtureTransfer

	accounts map[common.Address]chadAcc

	mu          sync.Mutex
	lastGenesis string

	deploys []*types.Transaction
}

// genesisOffset determines from which point we need generate contract addresses
func New(proxyAddress common.Address) *Chad {
	c := &Chad{
		proxyAddress: proxyAddress,

		transfers:   make([]*fixtureTransfer, 0),
		emissions:   make([]*fixtureEmission, 0),
		accsMap:     make(map[common.Address]*fixtureAcc),
		accList:     make([]*fixtureAcc, 0),
		emissionsId: make(map[common.Hash]*fixtureEmission),
		transfersId: make(map[common.Hash]*fixtureTransfer),
	}

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
		tx, sign := c.getTx(acc, payload, c.proxyAddress)
		var decTx = &types.Transaction{}
		b, err := rlp.EncodeToBytes(tx)
		if err != nil {
			log.Fatal(err)
		}
		err = rlp.DecodeBytes(b, decTx)
		if err != nil {
			log.Fatal(err)
		}

		emission := &fixtureEmission{
			emissionWallet: acc.WalletId,
			transaction:    tx,
			signature:      sign,
		}
		c.emissions = append(c.emissions, emission)
		c.emissionsId[tx.Id()] = emission
	}
}

func preCheck(tx *types.Transaction) (bool, error) {
	sigPublicKey, err := crypto.Ecrecover(tx.Hash().Bytes(), tx.Signature())
	if err != nil {
		return false, err
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(sigPublicKey[1:])[12:])
	return bytes.Equal(addr.Bytes(), tx.From().Bytes()), nil
}

func (c *Chad) InitTransfers(transfersAmount int) {
	for i := 0; i < transfersAmount; i++ {
		acc1, acc2 := c.getRandomAccsPair()
		payload := packTX(methodTransfer, acc1.WalletId, acc2.WalletId, big.NewInt(transferTokens))
		tx, sign := c.getTx(acc1, payload, c.proxyAddress)
		transfer := &fixtureTransfer{
			fromAccount: acc1,
			toAccount:   acc2,
			transaction: tx,
			signature:   sign,
		}

		c.transfers = append(c.transfers, transfer)
		c.transfersId[tx.Id()] = transfer
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

func (c *Chad) getTx(acc *fixtureAcc, payload []byte, to common.Address) (*types.Transaction, []byte) {
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
	return genTx, sign
}

func (c *Chad) getRandomAccsPair() (*fixtureAcc, *fixtureAcc) {
	first := rand.Intn(len(c.accList))
	second := first
	for second == first {
		second = rand.Intn(len(c.accList))
	}
	return c.accList[first], c.accList[second]
}

func (c *Chad) InitDeploys() {
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
