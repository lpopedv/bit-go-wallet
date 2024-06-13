package pkg

import (
	"crypto/sha256"
)

func GetCheckSum(input []byte) []byte {
	firstHash := sha256.New()
	firstHash.Write(input)

	first := firstHash.Sum(nil)

	secondHash := sha256.New()
	secondHash.Write(first)

	second := secondHash.Sum(nil)

  checkSum := second[:4]
  
  return checkSum
}
