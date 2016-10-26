# zgenerate

Offline BIP32 HD wallet generator for ZCash.

Currently returns the first address associated with m/0'/0/0 (hardened key for account 0/external account)

##Pre-requisites
* Golang 1.7.3 (lower versions may work but this is what I developed with)
* Git

##Build
go get -u github.com/btcsuite/btcutil
go get -u github.com/blackkeyboard/mneumonic

##Usage
To generate a wallet:
	
~~~~
zgenerate [-t]

Options
-t generate testnet addresses
~~~~

To retrieve the first address from the HD wallet:
	
~~~~
zretrieve [-t] -passphrase=<passphrase>

Options
-t generate testnet addresses	
~~~~

To import the private key into ZCash:
~~~~
./zcash-cli importprivkey "private_key_from_zgenerate"
~~~~
