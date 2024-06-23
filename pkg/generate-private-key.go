package pkg

import (
	"crypto/rand"
)

func GeneratePrivateKey() ([]byte, error) {
	privateKeyBytes := make([]byte, 32)

	_, err := rand.Read(privateKeyBytes)

	if err != nil {
		return nil, err
	}

	return privateKeyBytes, nil
}
