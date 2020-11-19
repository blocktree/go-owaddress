package test

import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_hns_AddressVerify_Valid(t *testing.T) {

	coin := "hns"
	expect := true

	address := "hs1qd3jn3mwhry65k8cq2semgxdwrqlg9mmt862c2k"


	valid, err := owaddress.Verify(coin, address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}


func Test_hns_AddressVerify_InValid(t *testing.T) {

	coin := "hns"
	expect := false

	address := "hs1qd3jn3mwhry65k8cq2semgxdwrqlg9mmt862cwk"


	valid, err := owaddress.Verify(coin, address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}
}
