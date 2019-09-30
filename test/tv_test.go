package test


import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_tv_AddressVerify_Valid(t *testing.T) {

	coin := "tv"
	expect := true

	Address := "tv3sAhFpc5mDwKW6CQo6byvZhhBLptPwKeH"


	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}


func Test_tv_AddressVerify_InValid(t *testing.T) {

	coin := "tv"
	expect := false

	Address := "tv3sAhFpc5mDwKW6CQo6byvZhhBLptPwKee"


	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}

}


