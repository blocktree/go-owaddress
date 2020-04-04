package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_nas_AddressVerify_Valid(t *testing.T) {

	coin := "nas"
	expect := true

	accountAddress := "n1TV3sU6jyzR4rJ1D7jCAmtVGSntJagXZHC"
	contractAddress := "n1sLnoc7j57YfzAVP8tJ3yK5a2i56QrTDdK"

	valid, err := owaddress.Verify(coin, accountAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify account valid address")
	}

	valid, err = owaddress.Verify(coin, contractAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify contract valid address")
	}
}

func Test_nas_AddressVerify_InValid(t *testing.T) {

	coin := "nas"
	expect := false

	accountAddress := "n1TV3sU6jyzR4rJ1D7jCAmtVGSntJagXZCC"
	contractAddress := "n1sLnoc7j57YfzAVP8tJ3yK5a2i56QrTDdD"

	valid, err := owaddress.Verify(coin, accountAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify account invalid address")
	}

	valid, err = owaddress.Verify(coin, contractAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify contract invalid address")
	}
}
