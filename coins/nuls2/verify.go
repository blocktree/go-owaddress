package nuls2

import (
	"bytes"
	"github.com/star001007/go-owaddress/address"
	"math/big"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

var (
	DefaultStruct = &AddressVerify{}
	CoinName      = "nuls2"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid(address string) bool {

	if address == "" {
		return false
	}

	if len(address) < 5 {
		return false
	}
	if address[:5] != "NULSd" {
		return false
	}

	address = address[5:]
	check := Base58Decode([]byte(address))
	return checkXOR(check)
}

func checkXOR(hashs []byte) bool {
	if len(hashs) < 24 {
		return false
	}
	body := hashs[:23]

	xor := byte(0)
	for _, v := range body {
		xor ^= v
	}
	if xor != hashs[23] {
		return false
	}
	return true

}

func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0
	for _, b := range input {
		if b == '1' {
			zeroBytes++
		} else {
			break
		}
	}

	payload := input[zeroBytes:]

	for _, b := range payload {
		charIndex := bytes.IndexByte(b58Alphabet, b)     //反推出余数
		result.Mul(result, big.NewInt(58))               //之前的结果乘以58
		result.Add(result, big.NewInt(int64(charIndex))) //加上这个余数

	}

	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{0x00}, zeroBytes), decoded...)

	return decoded[:24]
}
