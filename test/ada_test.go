package test

import (
	"github.com/star001007/go-owaddress"
	"github.com/star001007/go-owaddress/coins/ada"
	"testing"
)

func Test_ada_AddressVerify_Valid(t *testing.T) {

	coin := ada.CoinName
	expect := true

	Address := "addr1w9jx45flh83z6wuqypyash54mszwmdj8r64fydafxtfc6jgrw4rm3"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_ada_AddressVerify_InValid(t *testing.T) {
	coin := ada.CoinName
	expect := false
	Address := ""
	valid, err := owaddress.Verify(coin, Address)
	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
