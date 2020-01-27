package nhss

import (
	"github.com/blocktree/go-owaddress/address"
	"github.com/blocktree/go-owaddress/utils"
	"github.com/blocktree/go-owcrypt"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName = "nhss"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {
	var(
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

		P2PKHPrefix = byte(0x3f)
		P2SHPrefix = byte(0x0a)
	)

	if address == "" {
		return false
	}

	decodeBytes, err := utils.Base58Decode(address, base58Alphabet)
	if err != nil || len(decodeBytes) != 25 {
		return false
	}

	if decodeBytes[0] != P2PKHPrefix && decodeBytes[0] != P2SHPrefix {
		return false
	}

	check := owcrypt.Hash(decodeBytes[:21], 0, owcrypt.HASH_ALG_DOUBLE_SHA256)[:4]

	for i := 0; i < 4; i ++ {
		if check[i] != decodeBytes[21 + i] {
			return false
		}
	}

	return true
}

