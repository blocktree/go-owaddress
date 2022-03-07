package test

import (
	"github.com/star001007/go-owaddress"
	"github.com/star001007/go-owaddress/coins/sol"
	"testing"
)

func Test_sol_AddressVerify_Valid(t *testing.T) {

	coin := sol.CoinName
	expect := true

	Address := "9Z7E42k46kxnBjAh8YGXDw3rRGwwxQUBYM7Ccrmwg6ZP"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_sol_AddressVerify_InValid(t *testing.T) {
	coin := sol.CoinName
	expect := false
	Address := "019Z7E42k46kxnBjAh8YGXDw3rRGwwxQUBYM7Ccrmwg6ZP0"
	valid, err := owaddress.Verify(coin, Address)
	t.Logf("valid:%v, err:%v", valid, err)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
