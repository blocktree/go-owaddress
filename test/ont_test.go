package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_ont_AddressVerify_Valid(t *testing.T) {

	coin := "ont"
	expect := true

	Address := "ATfZt5HAHrx3Xmio3Ak9rr23SyvmgNVJqU"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}

func Test_ont_AddressVerify_InValid(t *testing.T) {

	coin := "ont"
	expect := false

	Address := "ATfZt5HAHrx3Xmio3Ak9rr23SyvmgNVJqq"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}
