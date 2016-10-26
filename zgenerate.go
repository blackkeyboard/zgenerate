package main

import (
	"flag"
	"log"

	"github.com/blackkeyboard/zgenerate/zcashcrypto"
)

func main() {
	var networkId zcashcrypto.NetworkId
	boolPtr := flag.Bool("test", false, "generate a testnet wallet")
	flag.Parse()

	// 2 bytes as per https://github.com/zcash/zcash/blob/master/src/chainparams.cpp
	if *boolPtr == true {
		networkId = zcashcrypto.NetworkId{0x1D, 0x25} //testnet
	} else {
		networkId = zcashcrypto.NetworkId{0x1C, 0xB8}
	}

	// Generate the wallet
	// More than 1 address can be generated but we'll pull off the first one
	wallet, err := zcashcrypto.CreateWallet(networkId, 1)

	if err != nil {
		log.Panicln(err.Error())
	}

	log.Println("Wallet generated!")
	log.Printf("Passphrase: %s\n", wallet.Passphrase)
	log.Printf("Address: %s\n", wallet.Addresses[0].Value)
	log.Printf("Privatekey: %s\n", wallet.Addresses[0].PrivateKey)
}
