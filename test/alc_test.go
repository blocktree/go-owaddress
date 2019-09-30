package test


import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_alc_AddressVerify_Valid(t *testing.T) {

	coin := "alc"
	expect := true

	p2pkhAddress := "AK84UEYQDyW2jXJ5yGUPdXS1DSeDicWg4o"
	//p2shAddress := "3BYx8ciMdywxd2bbn5h9V7EAZtzLg2RhhX"
	//bech32Address := "bc1qvgclzqz7smqr6haag9mknpwsjnxtdqkncr64kd"


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


func Test_alc_AddressVerify_InValid(t *testing.T) {

	coin := "alc"
	expect := false

	p2pkhAddress := "AK84UEYQDyW2jXJ5yGUPdXS1DSeDicWg45"
	//p2shAddress := "3BYx8ciMdywxd2bbn5h9V7EAZtzLg2R0hX"
	//bech32Address := "bc1qvgclzqz7smqr6haag9mknpwsjn3tdqkncr64kd"


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

