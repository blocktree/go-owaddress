package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_ltc_AddressVerify_Valid(t *testing.T) {

	expect := true
	coin := "ltc"

	p2pkhAddress := "LiZY38njUbKK5kMTzrdUA8Cq9K8D9HBwKm"
	p2shAddress := "31nM1cyzC3q8i8AHPK8QDFkLV8ecnuuUCG"
	bech32Address := "ltc1qqqqqpuwrmhu2k6vr97gsq4lfntyv7w7xxy2nal"
	p2sh2Address := "MKhcngTM1pk9a5z4jb58ZW3t7xFhvMDiDc"

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

	valid, err = owaddress.Verify(coin, bech32Address)

	if err != nil {
		t.Error(err)
	}

	valid, err = owaddress.Verify(coin, p2sh2Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify Bech32 valid address")
	}
}

func Test_ltc_AddressVerify_InValid(t *testing.T) {

	coin := "ltc"
	expect := false

	p2pkhAddress := "LiZY38njUbKK5kMTzrdUA8Cq9K8D9HBwKp"
	p2shAddress := "31nM1cyzC3q8i8AHPK8QDFkLV8ecnuuUC4"
	bech32Address := "ltc1qqqqqpuwrmhu2k6vr97gsq4lfntyv7w7xxy2na7"

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

	valid, err = owaddress.Verify(coin, bech32Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify Bech32 invalid address")
	}
}
