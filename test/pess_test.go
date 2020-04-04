package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_pess_AddressVerify_Valid(t *testing.T) {

	coin := "pess"
	expect := true

	Address := "GBYSQ2RMGCMW22FIGYVES7ZA6BCBM7ZYXCSRNXL6OPCYD7C4RTUOQZ4Y"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}

}

func Test_pess_AddressVerify_InValid(t *testing.T) {

	coin := "pess"
	expect := false

	Address := "GBYSQ2RMGCMW22FIGYVES7ZA6BCBM7ZYXCSRNXL6OPCYD7C4RTUOQZ3Y"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}

}
