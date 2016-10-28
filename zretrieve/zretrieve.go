// retrieves the addresses and priv keys associated with a mneumonic
package main

import (
	"flag"
	"log"

	"github.com/blackkeyboard/zgenerate/zcashcrypto"
)

func main() {
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

	log.Println("Wallet retrieved")
	log.Printf("Passphrase: %s\n", passphrase)
	log.Printf("Address\t\t\t\tPrivate key")

	for i := 0; i <= numAddresses-1; i++ {
		wallet, err := zcashcrypto.GetWalletFromPassphrase(!test, passphrase, uint32(i))

		if err != nil {
			log.Panicln(err.Error())
		}
		log.Printf("%s\t%s\n", wallet.Addresses[0].Value, wallet.Addresses[0].PrivateKey)
	}
}
