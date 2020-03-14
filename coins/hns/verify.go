package hns

import (
	"github.com/blocktree/go-owaddress/address"
	"github.com/blocktree/go-owaddress/utils"
	"strings"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName = "hns"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {
	var(
		bech32Alphabet = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"
		Bech32Prefix = "hs"
	)

	if address == "" {
		return false
	}

	if strings.Index(address, Bech32Prefix+"1") == 0 { // bech32 address
		hash, err := utils.Bech32Decode(address, bech32Alphabet)
		if err != nil || len(hash) != 20 {
			return false
		}

		return true
	}

	return false
}

