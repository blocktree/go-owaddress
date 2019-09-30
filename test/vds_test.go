package test


import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_vds_AddressVerify_Valid(t *testing.T) {

	coin := "vds"
	expect := true

	p2pkhAddress := "VcimwAzkqBDdFhzTunS2hQFrzBuXGmBHYYB"
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
}


func Test_vds_AddressVerify_InValid(t *testing.T) {

	coin := "vds"
	expect := false

	p2pkhAddress := "VcimwAzkqBDdFhzTunS2hQFrzBuXGmBHYYY"
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


}
