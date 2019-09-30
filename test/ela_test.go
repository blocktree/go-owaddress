package test


import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_ela_AddressVerify_Valid(t *testing.T) {

	coin := "ela"
	expect := true

	Address := "Eb1r8zaS3qbsRFH4j4GADshJCqFZ84ZM8u"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}


func Test_ela_AddressVerify_InValid(t *testing.T) {

	coin := "ela"
	expect := false

	Address := "Eb1r8zaS3qbsRFH4j4GADshJCqFZ84ZM6u"

	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}
}
