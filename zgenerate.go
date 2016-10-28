package main

import (
	"flag"
	"log"

	"github.com/blackkeyboard/zgenerate/zcashcrypto"
)

func main() {
	//	var networkId zcashcrypto.NetworkId
	boolPtr := flag.Bool("test", false, "generate a testnet wallet")
	nPtr := flag.Int("n", 1, "Number of addresses to generate")
	flag.Parse()

	// Generate the wallet
	wallet, err := zcashcrypto.CreateWallet(!(*boolPtr), *nPtr)

	if err != nil {
		log.Panicln(err.Error())
	}

	log.Println("Wallet generated!")
	log.Printf("Passphrase: %s\n", wallet.Passphrase)
	log.Printf("Address\t\t\t\tPrivate key")

	for i := 0; i <= len(wallet.Addresses)-1; i++ {
		log.Printf("%s\t%s\n", wallet.Addresses[i].Value, wallet.Addresses[i].PrivateKey)
	}
}
