package eos

import (
	"regexp"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "eos"

	reg = regexp.MustCompile("^[1-5a-z]{1,12}$")
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {
	return reg.MatchString(address)
}
