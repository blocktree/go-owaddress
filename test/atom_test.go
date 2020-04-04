package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_atom_AddressVerify_Valid(t *testing.T) {

	coin := "atom"
	expect := true

	Address := "cosmos1xv66sa5tlplm68j4fec6stdzszg3pcvswag06j"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}

func Test_atom_AddressVerify_InValid(t *testing.T) {

	coin := "atom"
	expect := false
	Address := "cosmos1xv66sa5tlplm68j4fec6stdzszg3pcvswag06z"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}
}
