package test


import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_moac_AddressVerify_Valid(t *testing.T) {

	coin := "moac"
	expect := true

	Address := "0x1dcbc4eac58965d9d845442df859a2f5434fec7a"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}

}

func Test_moac_AddressVerify_InValid(t *testing.T) {

	coin := "moac"
	expect := false

	Address := "0x1dcbc4eac58965d9d845442df859a2f5434fec7t"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}

}
