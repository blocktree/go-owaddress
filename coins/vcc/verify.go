package vcc

import (
	"encoding/hex"
	"strings"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "vcc"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {

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
