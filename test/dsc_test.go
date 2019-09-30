package test


import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_dsc_AddressVerify_Valid(t *testing.T) {

	coin := "dsc"
	expect := true

	p2pkhAddress := "D9JbM5dMeJ4QhmrvaANawXuNYNKgWXDGds"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}

}


func Test_dsc_AddressVerify_InValid(t *testing.T) {

	coin := "dsc"
	expect := false

	p2pkhAddress := "D9JbM5dMeJ4QhmrvaANawXuNYNKgWXDGdA"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}

}
