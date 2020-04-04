package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_vlx_AddressVerify_Valid(t *testing.T) {

	coin := "vlx"
	expect := true

	p2pkhAddress := "VLY5wrUfsjmKvTJGXEWwNpXQ6Wkzy8HU7Eb"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}

}

func Test_vlx_AddressVerify_InValid(t *testing.T) {

	coin := "vlx"
	expect := false

	p2pkhAddress := "VLY5wrUfsjmKvTJGXEWwNpXQ6Wkzy8HU7EE"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}

}
