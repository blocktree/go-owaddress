package sol

import (
	"regexp"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "solana"

	reg = regexp.MustCompile("^[1-9A-HJ-NP-Za-km-z]{32,44}$")
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {
	return reg.MatchString(address)
}
