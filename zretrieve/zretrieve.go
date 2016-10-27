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
	nPtr := flag.Int("n", 1, "Number of addresses to retrieve")

	flag.Parse()
	var passphrase string = *strPtr
	var test bool = *boolPtr
	var numAddresses int = *nPtr

	if passphrase == "" {
		log.Fatalln("Passphrase must be specified")
	}

	// 2 bytes as per https://github.com/zcash/zcash/blob/master/src/chainparams.cpp
	if test == true {
		networkId = zcashcrypto.NetworkId{0x1D, 0x25} //testnet
	} else {
		networkId = zcashcrypto.NetworkId{0x1C, 0xB8}
	}

	log.Println("Wallet retrieved")
	log.Printf("Passphrase: %s\n", passphrase)
	log.Printf("Address\t\t\t\tPrivate key")

	for i := 0; i <= numAddresses-1; i++ {
		wallet, err := zcashcrypto.GetWalletFromPassphrase(passphrase, networkId, uint32(i))

		if err != nil {
			log.Panicln(err.Error())
		}
		log.Printf("%s\t%s\n", wallet.Addresses[0].Value, wallet.Addresses[0].PrivateKey)
	}
}
