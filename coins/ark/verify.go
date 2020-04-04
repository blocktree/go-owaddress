package ark

import (
	"github.com/star001007/go-owaddress/address"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "ark"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {

	return true
}
