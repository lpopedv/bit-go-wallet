package main

import (
	"encoding/hex"
	"fmt"

	"github.com/lpopedv/bit-go-wallet/pkg"
)

func main() {
	privateKey, err := pkg.GeneratePrivateKey()

	if err != nil {
		panic(err)
	}

	wif := pkg.GenerateWif([]byte{0x80}, privateKey)
	publicKey := pkg.GeneratePublicKey(privateKey)
	privateKeyHex := hex.EncodeToString(privateKey)

btcWallet := pkg.GenerateBitcoinWallet([]byte{0x00}, publicKey)

fmt.Printf("Private Key: %s \nWIF: %s \nPublic Key: %s \nBitcoin Wallet: %s", privateKeyHex, wif, publicKey, btcWallet)
}
