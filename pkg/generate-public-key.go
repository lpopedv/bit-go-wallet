package pkg

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec"
)

func GeneratePublicKey(privateKey []byte) string {
	_, publicKey := btcec.PrivKeyFromBytes(btcec.S256(), privateKey)

	publicKeyBytes := publicKey.SerializeUncompressed()

	return string(hex.EncodeToString(publicKeyBytes))
}
