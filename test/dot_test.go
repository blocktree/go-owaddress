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
	"fmt"
	"github.com/blocktree/go-owaddress"
	"testing"
)

func Test_ksm_ed25519_AddressVerify_Valid(t *testing.T) {
	coin := "ksm"

	addressArr := make([]string, 0)
	addressArr = append(addressArr, "HcKJ5nYJoTXuKTKv96mdLbmS7MYwLZYRbUgLTXB4D8VRFLj")	//正确
	addressArr = append(addressArr, "HcKJ5nYJoTXuKTKv96mdLbmS7MYwLZYRbUgLTXB4D8VRFLi")	//改了最后一位
	addressArr = append(addressArr, "AcKJ5nYJoTXuKTKv96mdLbmS7MYwLZYRbUgLTXB4D8VRFLj")	//改了第一位
	addressArr = append(addressArr, "HcKJ5nYJoTXuKTKv96mdLbmS7MYwLZ26SUgLTXB4D8VRFLj")	//改了中间

	for i := 0; i < len(addressArr); i++ {
		address := addressArr[i]
		valid, err := owaddress.Verify(coin, address)

		if err != nil {
			t.Error(err)
		}

		fmt.Println(address, " isvalid : ", valid)
	}
}

func Test_ksm_sr25519_AddressVerify_Valid(t *testing.T) {
	coin := "ksm"

	addressArr := make([]string, 0)
	addressArr = append(addressArr, "CreEvbSZbf7bJ6ymX6UKNWmwBxmyXK6j29QWQGVvtr7TLsn")	//正确
	addressArr = append(addressArr, "CreEvbSZbf7bJ6ymX6UKNWmwBxmyXK6j29QWQGVvtr7TLs7")	//改了最后一位
	addressArr = append(addressArr, "WreEvbSZbf7bJ6ymX6UKNWmwBxmyXK6j29QWQGVvtr7TLsn")	//改了第一位
	addressArr = append(addressArr, "WreEvbSZbf7bJ6ymX6UKNWm1516yXK6j29QWQGVvtr7TLsn")	//改了中间

	for i := 0; i < len(addressArr); i++ {
		address := addressArr[i]
		valid, err := owaddress.Verify(coin, address)

		if err != nil {
			t.Error(err)
		}

		fmt.Println(address, " isvalid : ", valid)
	}
}
