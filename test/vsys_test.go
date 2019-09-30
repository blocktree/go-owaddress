package test


import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_vsys_AddressVerify_Valid(t *testing.T) {

	coin := "vsys"
	expect := true

	Address := "ARQEGuxzau9ZSsPgWWHNJYgVPUxJYQeGb4F"


	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}


func Test_vsys_AddressVerify_InValid(t *testing.T) {

	coin := "vsys"
	expect := false

	Address := "ARQEGuxzau9ZSsPgWWHNJYgVPUxJYQeGb4f"


	valid, err := owaddress.Verify(coin, Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
	}
}

