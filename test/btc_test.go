package test

import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func TestAddressVerify_Valid(t *testing.T) {

	expect := true

	p2pkhAddress := "19xD3nnvEiu7Uqd8irRvF3j5ExLb4ZtSju"
	p2shAddress := "3BYx8ciMdywxd2bbn5h9V7EAZtzLg2RhhX"
	bech32Address := "bc1qvgclzqz7smqr6haag9mknpwsjnxtdqkncr64kd"


	valid, err := owaddress.Verify("btc", p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid addredd")
	}


	valid, err = owaddress.Verify("btc", p2shAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH valid addredd")
	}

	valid, err = owaddress.Verify("btc", bech32Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify Bech32 valid addredd")
	}
}


func TestAddressVerify_InValid(t *testing.T) {

	expect := false

	p2pkhAddress := "19xD3nnvEiu7Uqd8irRvF3j5ExLb4ZtSj2"
	p2shAddress := "3BYx8ciMdywxd2bbn5h9V7EAZtzLg2R0hX"
	bech32Address := "bc1qvgclzqz7smqr6haag9mknpwsjn3tdqkncr64kd"


	valid, err := owaddress.Verify("btc", p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid addredd")
	}


	valid, err = owaddress.Verify("btc", p2shAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH valid addredd")
	}

	valid, err = owaddress.Verify("btc", bech32Address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify Bech32 valid addredd")
	}
}