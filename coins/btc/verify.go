package btc

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "btc"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {
	_, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	if err != nil {
		return false
	}
	return true
}
