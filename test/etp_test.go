package test


import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_etp_AddressVerify_Valid(t *testing.T) {

	coin := "etp"
	expect := true

	p2pkhAddress := "MTDcfh43xT93odL1Y2uULhRLeWED2fDvBX"
	p2shAddress := "33WuUsfKDHGho1KNSTknixEzYriUr2do8K"


	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}


	valid, err = owaddress.Verify(coin, p2shAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH valid address")
	}
	
}


func Test_etp_AddressVerify_InValid(t *testing.T) {

	coin := "etp"
	expect := false

	p2pkhAddress := "znkvsEfqqiJ7r9MPiUnoH4bUdkBKDAGx3mm"
	//p2shAddress := "3BYx8ciMdywxd2bbn5h9V7EAZtzLg2R0hX"
	//bech32Address := "bc1qvgclzqz7smqr6haag9mknpwsjn3tdqkncr64kd"


	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}


	//valid, err = owaddress.Verify(coin, p2shAddress)
	//
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//if valid != expect {
	//	t.Error("Failed to verify P2SH invalid address")
	//}
	//
	//valid, err = owaddress.Verify(coin, bech32Address)
	//
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//if valid != expect {
	//	t.Error("Failed to verify Bech32 invalid address")
	//}
}