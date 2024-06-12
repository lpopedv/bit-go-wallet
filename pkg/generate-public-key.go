package pkg

import (
	"encoding/hex"

	"github.com/btcsuite/btcd/btcec"
)

func GeneratePublicKey(privateKey string) string {
	_, publicKey := btcec.PrivKeyFromBytes(btcec.S256(), []byte(privateKey))

  return hex.EncodeToString(publicKey.SerializeUncompressed())
}
