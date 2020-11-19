package dot

import (
	"github.com/blocktree/go-owaddress/address"
	"github.com/blocktree/go-owaddress/utils"
	"github.com/blocktree/go-owcrypt"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinNameDOT = "dot"
	CoinNameKSM = "ksm"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {
	var(
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

		P2PKHPrefix = byte(0x02)
		prefix = []byte{0x02}
		ssPrefix = []byte{0x53, 0x53, 0x35, 0x38, 0x50, 0x52, 0x45}
	)
	decodeBytes, err := utils.Base58Decode(address, base58Alphabet)
	if err != nil || len(decodeBytes) != 35 {
		return false
	}
	if decodeBytes[0] != P2PKHPrefix {
		return false
	}
	pub := decodeBytes[1: len(decodeBytes)-2 ]

	data := append(prefix, pub...)
	input := append(ssPrefix, data...)
	checkSum := owcrypt.Hash(input, 64, owcrypt.HASH_ALG_BLAKE2B)[:2]

	for i := 0; i < 2; i ++ {
		if checkSum[i] != decodeBytes[33 + i] {
			return false
		}
	}
	if len(pub) != 32 {
		return false
	}

	return true
}
