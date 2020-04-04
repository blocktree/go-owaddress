package btc

import (
	"github.com/blocktree/go-owaddress/address"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "btc"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {
	_, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	if err != nil {
		return false
	}
	return true
}
