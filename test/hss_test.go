package test


import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_hss_AddressVerify_Valid(t *testing.T) {

	coin := "hss"
	expect := true

	p2pkhAddress := "HNoit5A6uo8h6ueLjySh3iiYFELDg3AkVW"
	//p2shAddress := "3BYx8ciMdywxd2bbn5h9V7EAZtzLg2RhhX"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH valid address")
	}

	//
	//valid, err = owaddress.Verify(coin, p2shAddress)
	//
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//if valid != expect {
	//	t.Error("Failed to verify P2SH valid address")
	//}
	//
	//valid, err = owaddress.Verify(coin, bech32Address)
	//
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//if valid != expect {
	//	t.Error("Failed to verify Bech32 valid address")
	//}
}


func Test_hss_AddressVerify_InValid(t *testing.T) {

	coin := "hss"
	expect := false

	p2pkhAddress := "HNoit5A6uo8h6ueLjySh3iiYFELDg3AkVV"
	//p2shAddress := "3BYx8ciMdywxd2bbn5h9V7EAZtzLg2R0hX"

	valid, err := owaddress.Verify(coin, p2pkhAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH invalid address")
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
