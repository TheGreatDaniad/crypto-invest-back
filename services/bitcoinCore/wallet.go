package bitcoinCore

import (
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
)

func (s Service) CreateAddress() (string, error) {
	client, err := getBTCCoreClient()
	if err != nil {
		panic(err)
	}
	defer client.Shutdown()
	address, err := client.GetNewAddress("wal")
	if err != nil {
		return "", err
	}
	return address.String(), nil
}
func (s Service) LoadWallet(walletPath string) error {
	client, err := getBTCCoreClient()
	if err != nil {
		panic(err)
	}
	defer client.Shutdown()

	_, err = client.LoadWallet(walletPath)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
func (s Service) GetAddressBalance(addressStr string) (float64, error) {
	client, err := getBTCCoreClient()
	if err != nil {
		panic(err)
	}
	defer client.Shutdown()

	address, err := btcutil.DecodeAddress(addressStr, &chaincfg.TestNet3Params)
	if err != nil {
		return 0, err
	}
	balance, err := client.GetReceivedByAddress(address)
	 //TODO handle address not found error 
	return balance.ToBTC(), err
}

func (s Service) GetAddressTransactions(addressStr string) ([]btcjson.ListTransactionsResult, error) {
	client, err := getBTCCoreClient()
	if err != nil {
		panic(err)
	}
	defer client.Shutdown()

	// address, err := btcutil.DecodeAddress(addressStr, &chaincfg.TestNet3Params)
	// if err != nil {
	// 	return nil, err
	// }

	txs, err := client.ListTransactionsCount(addressStr,10)
	if err != nil {
		return nil, err
	}

	return txs, nil
}
func generateWalletSeed() (string, error) {
	// Generate a new random 256-bit private key.
	key, err := btcec.NewPrivateKey()
	if err != nil {
		return "", err
	}

	// Convert the private key to a hex-encoded string.
	return hex.EncodeToString(key.Serialize()), nil
}
