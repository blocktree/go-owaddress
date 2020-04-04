package xif

import (
	"encoding/hex"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinNameXIF   = "xif"
	CoinNameAUSD  = "ausd"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {
	pub, err := hex.DecodeString(address)
	if err != nil {
		return false
	}

	if len(pub) != 33 {
		return false
	}

	return true
}
