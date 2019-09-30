package bnb

import (
	"github.com/blocktree/go-owaddress/address"
	"github.com/blocktree/go-owaddress/utils"
	"strings"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName = "bnb"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {
	var(
		Alphabet = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"
		Prefix = "bnb"
	)

	if address == "" {
		return false
	}

	if strings.Index(address, Prefix) == 0 { // bech32 address
		hash, err := utils.Bech32Decode(address, Alphabet)
		if err != nil || len(hash) != 20 {
			return false
		}

		return true
	}

	return false
}

