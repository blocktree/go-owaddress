package moac

import (
	"strings"

	"github.com/star001007/go-owaddress/utils"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "moac"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {

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
