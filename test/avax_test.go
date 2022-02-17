package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_avax_AddressVerify_Valid(t *testing.T) {

	coin := "avax"
	expect := true

	Address := "0xfDE9d08efEF15aB39c55cA29FFD71eB224B5daab"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_avax_AddressVerify_InValid(t *testing.T) {

	coin := "avax"
	expect := false

	Address := "0x1fDE9d08efEF15aB39c55cA29FFD71eB224B5daab"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
