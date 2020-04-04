package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_rcp_AddressVerify_Valid(t *testing.T) {

	coin := "rcp"
	expect := true

	Address := "rHbtcv6GanWzbm4NMQGuj19BjL27YBiQNg"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}

func Test_rcp_AddressVerify_InValid(t *testing.T) {

	coin := "rcp"
	expect := false

	Address := "rHbtcv6GanWzbm4NMQGuj19BjL27YBiQNN"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}
}
