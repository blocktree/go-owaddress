/*
 * Copyright 2019 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package test

import (
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_xif_AddressVerify_Valid(t *testing.T) {

	coin := "xif"
	expect := true

	address := "028c175947c2e0e1d992245d59b3bab9933d124f30d4156e919156dc4762bd34b2"

	valid, err := owaddress.Verify(coin, address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify XIF valid address")
	}
}

func Test_ausd_AddressVerify_Valid(t *testing.T) {

	coin := "ausd"
	expect := true

	address := "028c175947c2e0e1d992245d59b3bab9933d124f30d4156e919156dc4762bd34b2"

	valid, err := owaddress.Verify(coin, address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify AUSD valid address")
	}
}

func Test_xif_AddressVerify_InValid(t *testing.T) {

	coin := "xif"
	expect := false

	address := "0x1dcbc4eac58965d9d845442df859a2f5434fec7a"

	valid, err := owaddress.Verify(coin, address)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify XIF invalid address")
	}
}
