package test


import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_trx_AddressVerify_Valid(t *testing.T) {

	coin := "trx"
	expect := true

	Address := "TAJTMJuzvAqB8wmdUjRBVJW8CozfgrhpX3"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}


func Test_trx_AddressVerify_InValid(t *testing.T) {

	coin := "trx"
	expect := false

	Address := "TAJTMJuzvAqB8wmdUjRBVJW8CozfgrhpX2"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}

}
