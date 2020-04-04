package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_sgu_AddressVerify_Valid(t *testing.T) {

	coin := "sgu"
	expect := true

	Address := "STB9Cs3cCwoQ11TidumhAdcmvhLTZfHFUg"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_sgu_AddressVerify_InValid(t *testing.T) {

	coin := "sgu"
	expect := false

	Address := "STB9Cs3cCwoQ11TidumhAdcmvhLTZfHFU1"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
