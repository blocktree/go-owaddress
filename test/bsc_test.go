package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_bsc_AddressVerify_Valid(t *testing.T) {

	coin := "bsc"
	expect := true

	Address := "0x488Bfd38071d681D7f14ee4F0c5180a76cF88bd3"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_bsc_AddressVerify_InValid(t *testing.T) {

	coin := "bsc"
	expect := false

	Address := "0x488Bfd38071d681D7f14ee4F0c5180a76cF88bd"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
