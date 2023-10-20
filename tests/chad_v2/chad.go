package chad_v2

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/Unheilbar/pebbke_wallets/binding"
	"github.com/Unheilbar/pebbke_wallets/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const genesisPrivate = "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
const emissionTokens = 10000000

type NodeClient interface {
	SendTransaction()
}

type chadAcc struct {
	Address common.Address
	Private *ecdsa.PrivateKey
}

type fixtureAcc struct {
	Address         common.Address
	Private         *ecdsa.PrivateKey
	WalletId        [32]byte
	ExpectedBalance *big.Int //expexcted balance gets calculated every time new transfer is generated
}

type fixtureTransfer struct {
	fromAccount *fixtureAcc
	toAccount   *fixtureAcc
	sendTime    time.Time
	transaction types.Transaction
}

type fixtureEmission struct {
	emissionWallet [32]byte
	transaction    types.Transaction
}

type Chad struct {
	proxyAddress common.Address
	accsMap      map[common.Address]*fixtureAcc
	accList      []*fixtureAcc
	transfers    []*fixtureTransfer

	accounts map[common.Address]chadAcc

	client      NodeClient
	mu          sync.Mutex
	lastGenesis string
}

// genesisOffset determines from which point we need generate contract addresses
func New(proxyAddress common.Address) *Chad {
	c := &Chad{
		proxyAddress: proxyAddress,

		transfers: make([]*fixtureTransfer, 0),
		accsMap:   make(map[common.Address]*fixtureAcc),
		accList:   make([]*fixtureAcc, 0),
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
	// for _, acc := range c.accList {

	// }
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

func (c *Chad) GenerateAccEmissionsTx(contrAddr common.Address) []*types.Transaction {
	var ret []*types.Transaction

	for _, acc := range c.accounts {
		payload := getContractEmissionPayload(acc.Address, acc.Address.Hash().Bytes()[:32], contrAddr)
		c.fakeBalances[acc.Address.Hex()] += emissionVal
		genTx := types.NewTx(types.TxData{
			From: acc.Address,
			To:   &contrAddr,
			Id:   crypto.Keccak256Hash([]byte{3, 2, 1}),
			Data: payload,
		})
		sign, err := crypto.Sign(genTx.Hash().Bytes(), acc.Private)
		if err != nil {
			log.Fatal("sign err", err)
		}
		genTx.WithSignature(sign)
		ret = append(ret, genTx)
	}
	return ret
}

func (c *Chad) GetTestAccByID(id uint) chadAcc {
	return c.orderedAccs[id]
}

func (c *Chad) GenerateTransfers(size int, contrAddr common.Address) []*types.Transaction {
	var ret []*types.Transaction
	var ptrFrom int
	var ptrTo int
	genesisId := 1
	for i := 0; i < size; i++ {
		accFrom := c.orderedAccs[ptrFrom%len(c.orderedAccs)]
		ptrFrom += i
		accTo := c.orderedAccs[ptrTo%len(c.orderedAccs)]
		ptrTo += 1
		data := getContractTransferPayload(accFrom.Address.Hash().Bytes()[:32], accTo.Address.Hash().Bytes()[:32], contrAddr)
		id := crypto.Keccak256Hash([]byte(fmt.Sprint(genesisId)))
		genesisId++

		genTx := types.NewTx(types.TxData{
			From: accFrom.Address,
			To:   &contrAddr,
			Id:   id,
			Data: data,
		})
		sign, err := crypto.Sign(genTx.Hash().Bytes(), accFrom.Private)
		if err != nil {
			log.Fatal("sign err", err)
		}
		genTx.WithSignature(sign)
		ret = append(ret, genTx)
		c.fakeBalances[accFrom.Address.Hex()] -= transferVal
		c.fakeBalances[accTo.Address.Hex()] += transferVal
	}

	return ret
}

func (c *Chad) GenerateAccs(size int) {
	c.accounts = make(map[common.Address]chadAcc)
	c.fakeBalances = make(map[string]uint)

	prev := genesisPrivate
	for i := 0; i < size; i++ {
		privateKey, err := crypto.HexToECDSA(prev)
		if err != nil {
			log.Fatal("err gen private", err)
		}

		newAddr := common.BytesToAddress(crypto.Keccak256(crypto.FromECDSAPub(&privateKey.PublicKey)[1:])[12:])
		c.accounts[newAddr] = chadAcc{Address: newAddr, Private: privateKey}
		c.orderedAccs = append(c.orderedAccs, c.accounts[newAddr])
		prev = string(crypto.Keccak256Hash([]byte(prev)).Hex()[2:])
	}
}

func (c *Chad) GetContractDeployTX(from common.Address, contrCode []byte, id common.Hash) *types.Transaction {
	genTx := types.NewTx(types.TxData{
		From: from,
		To:   nil,
		Id:   id,
		Data: contrCode,
	})
	return genTx
}

const emissionVal = 100000000

func getContractEmissionPayload(pubKey common.Address, wallet []byte, contrAddr common.Address) []byte {
	var arr [32]byte
	copy(arr[:], wallet)
	data, err := packTX("emission", pubKey, arr, big.NewInt(emissionVal))
	if err != nil {
		log.Fatal("err generate emission input", err, data)
	}

	return data
}

const transferVal = 1

func getContractTransferPayload(fromWallet []byte, toWallet []byte, contrAddr common.Address) []byte {
	var arrFrom [32]byte
	copy(arrFrom[:], fromWallet)
	var arrTo [32]byte
	copy(arrTo[:], toWallet)
	data, err := packTX("transfer", arrFrom, arrTo, big.NewInt(transferVal))
	if err != nil {
		log.Fatal("err generate transfer input", err, data)
	}

	return data
}

func packTX(method string, params ...interface{}) ([]byte, error) {
	bind := binding.ProxyABI

	pabi, err := abi.JSON(strings.NewReader(bind))

	if err != nil {
		return nil, err
	}

	input, err := pabi.Pack(method, params...)
	if err != nil {
		return nil, err
	}

	return input, nil
}
