package pkg

import (
	"crypto/rand"
	"encoding/hex"
)

func GeneratePrivateKey() (string, error) {
  privateKeyBytes := make([]byte, 32)

  _, err := rand.Read(privateKeyBytes)

  if err != nil {
    return "", err
  }

  privateKeyHex := hex.EncodeToString(privateKeyBytes)


  return privateKeyHex, nil
}
