package pkg

import (
	"github.com/btcsuite/btcutil/base58"
)

func GenerateWif(version, privateKey []byte) string {
	privateKeyWithVersion := append(version, privateKey...)

	checkSum := GetCheckSum(privateKeyWithVersion)

	wif := append(privateKeyWithVersion, checkSum...)

	base68Wif := base58.Encode(wif)

	return base68Wif
}
