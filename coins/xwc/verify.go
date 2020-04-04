package xwc

import (
	"github.com/blocktree/go-owcrypt"

	"github.com/star001007/go-owaddress/utils"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "xwc"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {
	var (
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

		XWCAddressPrefix_mainnet = "XWC"
	)

	if address == "" {
		return false
	}

	base := address[len(XWCAddressPrefix_mainnet):]

	decodeBytes, err := utils.Base58Decode(base, base58Alphabet)
	if err != nil || len(decodeBytes) != 25 {
		return false
	}

	check := owcrypt.Hash(decodeBytes[:21], 0, owcrypt.HASH_ALG_RIPEMD160)[:4]

	for i := 0; i < 4; i++ {
		if check[i] != decodeBytes[21+i] {
			return false
		}
	}

	return true
}
