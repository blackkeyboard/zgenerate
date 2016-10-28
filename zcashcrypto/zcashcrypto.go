// Standalone library to generate zcash addresses
package zcashcrypto

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/blackkeyboard/mneumonic"
	"github.com/blackkeyboard/zgenerate/base58"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/btcsuite/golangcrypto/ripemd160"
)

type ZcashWallet struct {
	Passphrase string         `json:"passphrase"`
	HexSeed    string         `json:"hexSeed"`
	Addresses  []ZcashAddress `json:"addresses"`
	RequestId  string         `json:"requestId"`
}

type ZcashAddress struct {
	Value      string `json:"value"`
	PublicKey  string `json:"publicKey"`
	PrivateKey string `json:"privateKey"`
}

func getExtendedKeyFromPassphrase(mainnet bool, passphrase string) (*hdkeychain.ExtendedKey, error) {
	var networkCfg chaincfg.Params

	// Switch depending on mainnet or testnet
	if mainnet == true {
		networkCfg = chaincfg.MainNetParams

	} else {
		networkCfg = chaincfg.TestNet3Params
	}

	m := mneumonic.FromWords(strings.Split(passphrase, " "))
	hexSeed := m.ToHex()

	hexValue, err := hex.DecodeString(hexSeed)

	if err != nil {
		return nil, err
	}

	masterKey, err := hdkeychain.NewMaster(hexValue, &networkCfg)
	if err != nil {
		return nil, err
	}

	// get m/0'/0/0
	// Hardened key for account 0. ie 0'
	acct0, err := masterKey.Child(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		return nil, err
	}

	// External account for 0'
	extAcct0, err := acct0.Child(0)
	if err != nil {
		return nil, err
	}

	return extAcct0, nil
}

func getAddressFromPassphrase(mainnet bool, passphrase string, position uint32) (ZcashAddress, error) {
	var returnValue ZcashAddress
	var networkId NetworkId
	var networkCfg chaincfg.Params

	// Switch depending on mainnet or testnet
	if mainnet == true {
		networkId = MainnnetId
		networkCfg = chaincfg.MainNetParams

	} else {
		networkId = TestnetId
		networkCfg = chaincfg.TestNet3Params
	}

	extendedKey, err := getExtendedKeyFromPassphrase(mainnet, passphrase)
	if err != nil {
		return returnValue, err
	}

	key, err := extendedKey.Child(uint32(position))
	if err != nil {
		return returnValue, err
	}

	// Serialize to compressed key bytes and pkhash
	pk, err := key.ECPubKey()
	if err != nil {
		return returnValue, err
	}
	pkSerialized := pk.SerializeCompressed()
	pkHash := btcutil.Hash160(pkSerialized)

	encodedAddress := base58.CheckEncode(pkHash[:ripemd160.Size], networkId)

	// Get the pubkey and serialise the compressed public key
	privKey, err := key.ECPrivKey()
	if err != nil {
		return returnValue, err
	}

	returnValue.Value = fmt.Sprintf("%s", encodedAddress)
	wif, err := btcutil.NewWIF(privKey, &networkCfg, true)

	if err != nil {
		return returnValue, err
	}

	returnValue.PrivateKey = wif.String()
	returnValue.PublicKey = hex.EncodeToString(privKey.PubKey().SerializeCompressed())

	return returnValue, nil
}

func CreateWallet(mainnet bool, numberOfAddressesToGenerate int) (ZcashWallet, error) {
	var wallet ZcashWallet
	var numAddresses int
	var networkId NetworkId
	var networkCfg chaincfg.Params

	// Switch depending on mainnet or testnet
	if mainnet == true {
		networkId = MainnnetId
		networkCfg = chaincfg.MainNetParams

	} else {
		networkId = TestnetId
		networkCfg = chaincfg.TestNet3Params
	}

	if numberOfAddressesToGenerate <= 0 {
		numAddresses = 20
	} else if numberOfAddressesToGenerate > 100 {
		numAddresses = 100
	} else {
		numAddresses = numberOfAddressesToGenerate
	}

	m := mneumonic.GenerateRandom(128)
	wallet.Passphrase = strings.Join(m.ToWords(), " ")
	wallet.HexSeed = m.ToHex()

	extendedKey, err := getExtendedKeyFromPassphrase(mainnet, wallet.Passphrase)
	if err != nil {
		return wallet, err
	}

	// Derive extended key (repeat this from 0 to number of addresses-1)
	for i := 0; i <= numAddresses-1; i++ {
		var address ZcashAddress

		key, err := extendedKey.Child(uint32(i))
		if err != nil {
			return wallet, err
		}

		// Serialize to compressed key bytes and pkhash
		pk, err := key.ECPubKey()
		if err != nil {
			return wallet, err
		}
		pkSerialized := pk.SerializeCompressed()
		pkHash := btcutil.Hash160(pkSerialized)

		encodedAddress := base58.CheckEncode(pkHash[:ripemd160.Size], networkId)

		// Get the pubkey and serialise the compressed public key
		privKey, err := key.ECPrivKey()
		if err != nil {
			return wallet, err
		}

		address.Value = fmt.Sprintf("%s", encodedAddress)
		wif, err := btcutil.NewWIF(privKey, &networkCfg, true)

		if err != nil {
			return wallet, err
		}

		address.PrivateKey = wif.String()
		address.PublicKey = hex.EncodeToString(privKey.PubKey().SerializeCompressed())

		wallet.Addresses = append(wallet.Addresses, address)
	}

	return wallet, nil
}

func GetWalletFromPassphrase(mainnet bool, passphrase string, position uint32) (ZcashWallet, error) {
	var result ZcashWallet
	var address ZcashAddress

	address, err := getAddressFromPassphrase(mainnet, passphrase, position)

	if err != nil {
		return result, err
	}

	result.Passphrase = passphrase
	result.Addresses = append(result.Addresses, address)

	return result, nil
}
