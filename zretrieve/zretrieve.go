// retrieves the addresses and priv keys associated with a mneumonic
package main

import (
	"flag"
	"log"

	"github.com/blackkeyboard/zgenerate/zcashcrypto"
)

func main() {
	var networkId zcashcrypto.NetworkId
	boolPtr := flag.Bool("test", false, "generate a testnet wallet")
	strPtr := flag.String("passphrase", "", "Passphrase for the wallet")

	flag.Parse()
	var passphrase string = *strPtr
	var test bool = *boolPtr

	if passphrase == "" {
		log.Fatalln("Passphrase must be specified")
	}

	// 2 bytes as per https://github.com/zcash/zcash/blob/master/src/chainparams.cpp
	if test == true {
		networkId = zcashcrypto.NetworkId{0x1D, 0x25} //testnet
	} else {
		networkId = zcashcrypto.NetworkId{0x1C, 0xB8}
	}

	wallet, err := zcashcrypto.GetWalletFromPassphrase(passphrase, networkId, 0)

	if err != nil {
		log.Panicln(err.Error())
	}

	log.Println("Wallet retrieved")
	log.Printf("Passphrase: %s\n", wallet.Passphrase)
	log.Printf("Address: %s\n", wallet.Addresses[0].Value)
	log.Printf("Privatekey: %s\n", wallet.Addresses[0].PrivateKey)

}
