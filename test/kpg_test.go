package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_kpg_AddressVerify_Valid(t *testing.T) {

	expect := true
	coin := "kpg"

	p2pkhAddress := "KDn2zSFQehXrK9vnU56PvZp7cHif9p9wTo"
	p2shAddress := "PZkfhQs1RUuPW7Rngrk3TM87earb9e2qxx"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}

	valid, err = owaddress.Verify(coin, p2shAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH valid address")
	}
}

func Test_kpg_AddressVerify_InValid(t *testing.T) {

	expect := false
	coin := "kpg"

	p2pkhAddress := "QjwHHRHUzaPebgmDkrtx3CxkchtDW5eB9e"
	p2shAddress := "M7zoS2sNa9jg5eYiMvaJeAR3pNJe8gCT12"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}

	valid, err = owaddress.Verify(coin, p2shAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH invalid address")
	}
}
