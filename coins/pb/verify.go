package pb

import (
	"strings"

	"github.com/star001007/go-owaddress/utils"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "pb"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {
	var (
		bech32Alphabet = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

		Bech32Prefix = "cosmos"
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

	return false
}
