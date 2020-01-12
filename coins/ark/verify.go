package ark

import (
	"github.com/blocktree/go-owaddress/address"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName = "ark"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {

	return true
}
