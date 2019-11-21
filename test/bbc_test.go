package test
import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_bbc_AddressVerify_Valid(t *testing.T) {

	coin := "bbc"
	expect := true

	Address := "11hdxdzd82h01esjb4th7963db2fwf90rp8cxjk5jsrgvjdf2rnx5tges"


	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify valid address")
	}
}


func Test_bbc_AddressVerify_InValid(t *testing.T) {

	coin := "bbc"
	expect := false
	Address := "11hdxdzd82h01esjb4th7963db2fwf90rp8cxjk5jsrgvjdf5rnx5tges"


	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify invalid address")
	}
}