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

func Test_ae_AddressVerify_Valid(t *testing.T) {

	coin := "ae"
	expect := true

	AEAddress := "ak_2QkttUgEyPixKzqXkJ4LX7ugbRjwCDWPBT4p4M2r8brjxUxUYd"

	valid, err := owaddress.Verify(coin, AEAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify AE valid address")
	}
}


func Test_ae_AddressVerify_InValid(t *testing.T) {

	coin := "ae"
	expect := false

	AEAddress := "0x1dcbc4eac58965d9d845442df859a2f5434fec7a"

	valid, err := owaddress.Verify(coin, AEAddress)

	if err != nil {
		t.Error(err)
	}

	if valid != expect {
		t.Error("Failed to verify AE invalid address")
	}
}
