package vlx

import (
	"github.com/blocktree/go-owcrypt"
	"github.com/star001007/go-owaddress/address"
	"github.com/star001007/go-owaddress/utils"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "vlx"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {
	var (
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

		P2PKHPrefix = []byte{15, 244}
	)

	if address == "" {
		return false
	}

	decodeBytes, err := utils.Base58Decode(address, base58Alphabet)
	if err != nil || len(decodeBytes) != 26 {
		return false
	}

	if decodeBytes[0] != P2PKHPrefix[0] || decodeBytes[1] != P2PKHPrefix[1] {
		return false
	}

	check := owcrypt.Hash(decodeBytes[:22], 0, owcrypt.HASH_ALG_DOUBLE_SHA256)[:4]

	for i := 0; i < 4; i++ {
		if check[i] != decodeBytes[22+i] {
			return false
		}
	}

	return true
}
