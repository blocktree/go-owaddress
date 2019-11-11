package xwc

import (
	"github.com/blocktree/go-owaddress/address"
	"github.com/blocktree/go-owaddress/utils"
	"github.com/blocktree/go-owcrypt"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName = "xwc"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {
	var(
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

	for i := 0; i < 4; i ++ {
		if check[i] != decodeBytes[21 + i] {
			return false
		}
	}

	return true
}


