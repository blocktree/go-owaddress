package test

import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_eva_AddressVerify_Valid(t *testing.T) {

	coin := "eva"
	expect := true

	Address := "eva1pn80qt83wzk9w4gs3muc8hw26cexlgav75mar0"


	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}


func Test_eva_AddressVerify_InValid(t *testing.T) {

	coin := "eva"
	expect := false
	Address := "eva1pn80qt83wzk9w4gs3muc8hw26cexlgav75marr"


	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}
}
