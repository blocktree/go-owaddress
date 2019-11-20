package test

import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_nuls2_AddressVerify_Valid(t *testing.T) {

	coin := "nuls2"
	expect := true

	Address := "NULSd6Hgap7WJw6inBDSccToUdpth8uXzPfvL"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}


func Test_nuls2_AddressVerify_InValid(t *testing.T) {

	coin := "nuls2"
	expect := false

	Address := "tNULSeBaNTcZo37gNC5mNjJuB39u8zT3TAy8jy"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}

