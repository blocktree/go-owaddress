package xvg

import (
	"github.com/blocktree/go-owcrypt"
	"github.com/star001007/go-owaddress/address"
	"github.com/star001007/go-owaddress/utils"
	"strings"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "xvg"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {
	var (
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
		bech32Alphabet = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

		P2PKHPrefix  = byte(0x1e)
		P2SHPrefix   = byte(0x21)
		Bech32Prefix = "xvg"
	)

	if address == "" {
		return false
	}

	if strings.Index(address, Bech32Prefix) == 0 { // bech32 address
		hash, err := utils.Bech32Decode(address, bech32Alphabet)
		if err != nil || len(hash) != 20 {
			return false
		}

		return true
	}

	decodeBytes, err := utils.Base58Decode(address, base58Alphabet)
	if err != nil || len(decodeBytes) != 25 {
		return false
	}

	if decodeBytes[0] != P2PKHPrefix && decodeBytes[0] != P2SHPrefix {
		return false
	}

	check := owcrypt.Hash(decodeBytes[:21], 0, owcrypt.HASH_ALG_DOUBLE_SHA256)[:4]

	for i := 0; i < 4; i++ {
		if check[i] != decodeBytes[21+i] {
			return false
		}
	}

	return true
}
