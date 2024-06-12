package pkg

import (
	"crypto/sha256"
	"encoding/hex"
)

func GetCheckSum(privateKeyHex string) string {
	privateKeyBytes := []byte(privateKeyHex)

	firstHash := sha256.New()
	firstHash.Write(privateKeyBytes)

	firstPrivateKeyHashed := firstHash.Sum(nil)

	secondHash := sha256.New()
	secondHash.Write(firstPrivateKeyHashed)

	secoundPrivateKeyHashed := secondHash.Sum(nil)

  checkSum := hex.EncodeToString(secoundPrivateKeyHashed[:4])
  
  return checkSum
}
