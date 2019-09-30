package test


import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_fiii_AddressVerify_Valid(t *testing.T) {

	coin := "fiii"
	expect := true

	p2pkhAddress := "fiiimMsC3cLHr8qxNooVR3zgGhYmprhuiabTX7"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}

}


func Test_fiii_AddressVerify_InValid(t *testing.T) {

	coin := "fiii"
	expect := false

	p2pkhAddress := "fiiimMsC3cLHr8qxNooVR3zgGhYmprhuiabTX8"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}
}
