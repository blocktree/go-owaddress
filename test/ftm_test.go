package test

import (
	"github.com/star001007/go-owaddress"
	"github.com/star001007/go-owaddress/coins/ftm"
	"testing"
)

func Test_ftm_AddressVerify_Valid(t *testing.T) {
	coin := ftm.CoinName
	expect := true

	Address := "0x21db06D44b57389722Ae318561B0b2E8828B938A"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_ftm_AddressVerify_InValid(t *testing.T) {
	coin := ftm.CoinName
	expect := false

	Address := "0x121db06D44b57389722Ae318561B0b2E8828B938A"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
