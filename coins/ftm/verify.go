package ftm

import (
	"encoding/hex"
	"strings"
)

const addressLength = 20

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "fantom"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {
	if address == "" {
		return false
	}

	if strings.Index(address, "0x") != 0 {
		return false
	}

	addressBytes, err := hex.DecodeString(address[2:])
	if err != nil {
		return false
	}

	if len(addressBytes) != addressLength {
		return false
	}

	return true
}
