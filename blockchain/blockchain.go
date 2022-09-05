package blockchain

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	//"go/types"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func InitNetwork(networkRPC string) (*ethclient.Client) {
	var err error
	client, err = ethclient.Dial(networkRPC)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to blockchain network")
	return client
}

func CreateWallet() (common.Address, *ecdsa.PrivateKey) {
	generatedPrivateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	//privateKeyBytes := crypto.FromECDSA(generatedPrivateKey)
	//privateKey := hexutil.Encode(privateKeyBytes)[2:]
	//fmt.Println(privateKey)
	publicKey := generatedPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, generatedPrivateKey
}

func ImportWallet(privateKey string) (common.Address, *ecdsa.PrivateKey) {
	importedPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := importedPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, importedPrivateKey
}

func NewTransactOpts(context context.Context , privateKey *ecdsa.PrivateKey) (*bind.TransactOpts){
	//fetch networkID
	networkID, err := client.ChainID(context)
	if err != nil {
		log.Fatal(err)
	}
	txOps, err := bind.NewKeyedTransactorWithChainID(privateKey, networkID)
	if err != nil {
		log.Fatal(err)
	}
	return txOps
}

func SetTransactOpts(address common.Address, context context.Context, txOps *bind.TransactOpts, value *big.Int, gasLimit uint64, gasPrice *big.Int) {
	nonce, err := client.PendingNonceAt(context, address)
	if err != nil {
		log.Fatal(err)
	}
	txOps.Nonce = big.NewInt(int64(nonce))
	txOps.Value = value
	txOps.GasLimit = gasLimit
	// if given gasPrice is 0, suggested gasPrice will be used
	if gasPrice == big.NewInt(0) {
		txOps.GasPrice, err = client.SuggestGasPrice(context)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		txOps.GasPrice = gasPrice
	}
}

func NewCallOpts(pending bool , from common.Address, blockNumber *big.Int, context context.Context) (*bind.CallOpts){
	callOps := &bind.CallOpts{
		Pending: pending,
		From: from,
		BlockNumber: blockNumber,
		Context: context,
	}
	return callOps
}

func DeployContract(address common.Address, privateKey *ecdsa.PrivateKey, name string, symbol string, supply *big.Int) (common.Address, *Blockchain) {
	nonce, err := client.PendingNonceAt(context.Background(), address)
    if err != nil {
        log.Fatal(err)
    }

	gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }
	chainID, err := client.ChainID(context.Background())
	if err != nil {
        log.Fatal(err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
        log.Fatal(err)
    }
	auth.From = address
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

	contractAddress, _, instance, err := DeployBlockchain(auth, client, name, symbol, address, supply)
	if err != nil {
		log.Fatal(err)
	}

	return contractAddress, instance
}

func TransferContract(ctx context.Context, privateKey *ecdsa.PrivateKey, address common.Address, to common.Address, amount *big.Int) (common.Hash, error) {
	txOps := NewTransactOpts(ctx, privateKey)
	SetTransactOpts(address, ctx, txOps, amount, uint64(300000), big.NewInt(0))
	blockInstance, err := NewBlockchain(address, client)
	if err != nil {
		log.Fatal("blockchain instance error: ",err)
	}

	tx, err := blockInstance.Transfer(txOps, to, amount)

	return tx.Hash(), err
}

func GetBalance(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error){
	return client.BalanceAt(ctx, account, blockNumber)
}

func GetTokenBalance(ctx context.Context, address common.Address) (*big.Int, error){
	callOps := bind.CallOpts{
		Pending: true,
		From: address,
	}
	blockInstance, err := NewBlockchain(address, client)
	if err != nil {
		log.Fatal("blockchain instance error: ",err)
	}
	return blockInstance.BalanceOf(&callOps, address)
}
