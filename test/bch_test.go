package test

import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_bch_AddressVerify_Valid(t *testing.T) {

	coin := "bch"
	expect := true

	p2pkhaddr_legacy := "1ArAfjvwXBiYhWcirLRc9CsvhuSgVqRQbc"
	p2pkhaddr_cashaddr := "bitcoincash:qpkqt9dxyae8c4r32m99lm24aw09vctc8ulss5czme"
	p2pkhaddr_cashaddr_short := "qpkqt9dxyae8c4r32m99lm24aw09vctc8ulss5czme"
	p2pkhaddr_copay := "CSK4EnH1QEh5beX9Y5kXiiVxL2f6Knung5"

	p2shaddr_legacy := "3CXRBKkvzB1QPJz2AVJYavEBr228v6Hkqe"
	p2shaddr_cashaddr := "bitcoincash:ppmd3gkmjxedpexxxtvvx4vydasd8qceus83xaa7sn"
	p2shaddr_cashaddr_short := "ppmd3gkmjxedpexxxtvvx4vydasd8qceus83xaa7sn"
	p2shcopay := "HHMXe8C1qVE51Us42AxhZJkisg39khHwTy"

	valid, err := owaddress.Verify(coin, p2pkhaddr_legacy)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH legacy valid address")
	}

	valid, err = owaddress.Verify(coin, p2pkhaddr_cashaddr)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH cash valid address")
	}

	valid, err = owaddress.Verify(coin, p2pkhaddr_cashaddr_short)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH cash valid address")
	}

	valid, err = owaddress.Verify(coin, p2pkhaddr_copay)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH copay valid address")
	}


	valid, err = owaddress.Verify(coin, p2shaddr_legacy)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH legacy valid address")
	}

	valid, err = owaddress.Verify(coin, p2shaddr_cashaddr)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH cash valid address")
	}

	valid, err = owaddress.Verify(coin, p2shaddr_cashaddr_short)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH cash valid address")
	}

	valid, err = owaddress.Verify(coin, p2shcopay)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH copay valid address")
	}
}


func Test_bch_AddressVerify_InValid(t *testing.T) {

	coin := "bch"
	expect := false

	p2pkhaddr_legacy := "1ArAfjvwXBiYhWcirLRc9CsvhuSgVqRcbc"
	p2pkhaddr_cashaddr := "bitcoincash:qpkqt9dxyae8c4r32m99lm24aw00vctc8ulss5czme"
	p2pkhaddr_cashaddr_short := "qpkqt9dxyae8c4r32m99lm24aw09vrtc8ulss5czme"
	p2pkhaddr_copay := "CSK4EnH1QEh5beX9Y5kXiiVxx2f6Knung5"

	p2shaddr_legacy := "3CXRBKkvzB1QPJz2AVJYavEBr2x8v6Hkqe"
	p2shaddr_cashaddr := "bitcoincash:ppmd3gkmjxedpexxxtvvx4vydasd8qceus8xxaa7sn"
	p2shaddr_cashaddr_short := "ppmd3gkmjxedpexxxtvvx4vydasd8qceux83xaa7sn"
	p2shcopay := "HHMXe8C1qVE51Us42AxhZJkisg39khxwTy"

	valid, err := owaddress.Verify(coin, p2pkhaddr_legacy)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH legacy valid address")
	}

	valid, err = owaddress.Verify(coin, p2pkhaddr_cashaddr)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH cash valid address")
	}

	valid, err = owaddress.Verify(coin, p2pkhaddr_cashaddr_short)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH cash valid address")
	}

	valid, err = owaddress.Verify(coin, p2pkhaddr_copay)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2PKH copay valid address")
	}


	valid, err = owaddress.Verify(coin, p2shaddr_legacy)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH legacy valid address")
	}

	valid, err = owaddress.Verify(coin, p2shaddr_cashaddr)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH cash valid address")
	}

	valid, err = owaddress.Verify(coin, p2shaddr_cashaddr_short)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH cash valid address")
	}

	valid, err = owaddress.Verify(coin, p2shcopay)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify P2SH copay valid address")
	}
}

