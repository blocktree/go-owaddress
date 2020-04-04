package test

import (
	"github.com/star001007/go-owaddress"
	"testing"
)

func Test_xwc_AddressVerify_Valid(t *testing.T) {

	coin := "xwc"
	expect := true

	p2pkhAddress := "XWCNeVn7JSzGQwK1GGMgbk2V3jCbc6x1p7ZAo"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify XWC valid address")
	}

}

func Test_xwc_AddressVerify_InValid(t *testing.T) {

	coin := "xwc"
	expect := false

	p2pkhAddress := "znkvsEfqqiJ7r9MPiUnoH4bUdkBKDAGx3mm"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify XWC valid address")
	}

}
