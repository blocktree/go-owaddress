package moac

import (
	"github.com/star001007/go-owaddress/address"
	"github.com/star001007/go-owaddress/utils"
	"strings"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "moac"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {

	if address == "" || len(address) != 42 {
		return false
	}

	if strings.Index(address, "0x") != 0 { // bech32 address
		return false
	}

	if !utils.IsLowerHex(address[2:]) {
		return false
	}

	return true
}
