package fiii

import (
	"github.com/blocktree/go-owcrypt"

	"github.com/star001007/go-owaddress/utils"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "fiii"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {
	var (
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

		P2PKHPrefix = []byte{0x40, 0xE7, 0xE9, 0x15}
	)

	if address == "" {
		return false
	}

	decodeBytes, err := utils.Base58Decode(address, base58Alphabet)
	if err != nil || len(decodeBytes) != 28 {
		return false
	}

	for i, d := range decodeBytes[:4] {
		if d != P2PKHPrefix[i] {
			return false
		}
	}

	check := owcrypt.Hash(decodeBytes[:24], 0, owcrypt.HASH_ALG_DOUBLE_SHA256)[:4]

	for i := 0; i < 4; i++ {
		if check[i] != decodeBytes[24+i] {
			return false
		}
	}

	return true
}
