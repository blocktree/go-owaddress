package test

import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_wicc_AddressVerify_Valid(t *testing.T) {

	coin := "wicc"
	expect := true

	Address := "WbcMwwvLv7DxkjzH8pQvok1wdomuSLfEGU"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}


func Test_wicc_AddressVerify_InValid(t *testing.T) {

	coin := "wicc"
	expect := false

	Address := "WbcMwwvLv7DxkjzH8pQvok1wdomuSLfEGG"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}

}

