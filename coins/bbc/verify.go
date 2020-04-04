package bbc

import (
	"encoding/base32"
	"encoding/binary"
	"github.com/star001007/go-owaddress/address"
	"strings"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "bbc"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {
	var (
		base32Alphabet = "0123456789abcdefghjkmnpqrstvwxyz"
		prefix         = "1"
	)

	if address == "" {
		return false
	}

	if strings.Index(address, prefix) != 0 {
		return false
	}

	pubkey, err := base32.NewEncoding(base32Alphabet).DecodeString(address[1:])

	if err != nil || len(pubkey) != 32+3 {
		return false
	}

	tmp := [4]byte{}

	binary.BigEndian.PutUint32(tmp[:], crc24q(pubkey[:32]))

	for i := 0; i < 3; i++ {
		if tmp[i+1] != pubkey[32+i] {
			return false
		}
	}

	return true
}
