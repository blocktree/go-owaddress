package test


import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_bnb_AddressVerify_Valid(t *testing.T) {

	coin := "bnb"
	expect := true

	Address := "bnb1p89xmrejmqpwmkye4z2pwtrw49n2vykf3nax33"


	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}


func Test_bnb_AddressVerify_InValid(t *testing.T) {

	coin := "bnb"
	expect := false

	Address := "bnb1p89xmrejmqpwmkye4z2pwtrw49n2vykf3nax32"


	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}
}
