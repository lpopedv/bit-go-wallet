package pkg

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

func GenerateBitcoinWallet(networkVersion []byte, publicKey string) string {
	publicKeyBytes, err := hex.DecodeString(publicKey)

	if err != nil {
		panic(err)
	}

	sha256Hash := sha256.New()
	sha256Hash.Write(publicKeyBytes)
	publicKeyHashed := sha256Hash.Sum(nil)

	ripemd160Hash := ripemd160.New()
	ripemd160Hash.Write(publicKeyHashed)
	publicKeyHashInRipemd := ripemd160Hash.Sum(nil)

	hashWithNetworkVersion := append(networkVersion, publicKeyHashInRipemd...)

	checkSum := GetCheckSum(hashWithNetworkVersion)

	wallet := append(hashWithNetworkVersion, checkSum...)

	return base58.Encode(wallet)
}
