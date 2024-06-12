package pkg

import "github.com/btcsuite/btcutil/base58"

func GenerateWif(version, privateKey string) string {
  privateKeyWithVersion := version + privateKey
  checkSum := GetCheckSum(privateKeyWithVersion)

  wif := privateKeyWithVersion + checkSum

  wifInBase58 := base58.Encode([]byte(wif))

  return wifInBase58 
}
