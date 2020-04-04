package sgu

import (
	"fmt"
	"strings"
)

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "sgu"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {

	result, _ := ValidateAddress(address)

	return result
}

func AddressToBytes(address string) ([]byte, error) {
	bytes, err := Decode(address)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func ValidateAddress(address string) (bool, error) {
	bytes, err := AddressToBytes(address)

	if err != nil {
		return false, err
	}

	return Byte2Hex(63) == Hex2Byte(bytes[:1]), nil
}

func Byte2Hex(data byte) string {
	return fmt.Sprintf("%x", data)
}

func Hex2Byte(data []byte) string {
	return strings.ToLower(fmt.Sprintf("%X", data))
}
