package test

import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_pb_AddressVerify_Valid(t *testing.T) {

	coin := "pb"
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


func Test_pb_AddressVerify_InValid(t *testing.T) {

	coin := "pb"
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
