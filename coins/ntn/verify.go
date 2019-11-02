package ntn

import (
	"encoding/hex"
	"github.com/blocktree/go-owaddress/address"
	"strings"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName = "ntn"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {

	if address == "" {
		return false
	}

	if strings.Index(address, "0x") != 0 {
		return false
	}

	_, err := hex.DecodeString(address[2:])
	if err != nil {
		return false
	}

	return true
}

