package xif

import (
	"encoding/hex"
	"github.com/blocktree/go-owaddress/address"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinNameXIF = "xif"
	CoinNameAUSD = "ausd"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {
	pub, err := hex.DecodeString(address)
	if err != nil {
		return false
	}

	if len(pub) != 33 {
		return false
	}

	return true
}

