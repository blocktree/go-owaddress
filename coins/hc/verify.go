package hc

import (
	"github.com/star001007/go-owaddress/utils"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "hc"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {
	var (
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

		P2PKHPrefix = []byte{0x09, 0x7F}
		P2SHPrefix  = []byte{0x09, 0x5A}
	)

	if address == "" {
		return false
	}

	decodeBytes, err := utils.Base58Decode(address, base58Alphabet)
	if err != nil || len(decodeBytes) != 26 {
		return false
	}

	if (decodeBytes[0] != P2PKHPrefix[0] || decodeBytes[1] != P2PKHPrefix[1]) &&
		(decodeBytes[0] != P2SHPrefix[0] || decodeBytes[1] != P2SHPrefix[1]) {
		return false
	}

	check := utils.DoubleBlake256(decodeBytes[:22])[:4]

	for i := 0; i < 4; i++ {
		if check[i] != decodeBytes[22+i] {
			return false
		}
	}

	return true
}
