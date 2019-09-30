package nas


import (
	"github.com/blocktree/go-owaddress/address"
	"github.com/blocktree/go-owaddress/utils"
	"github.com/blocktree/go-owcrypt"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName = "nas"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {
	var(
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

		accountPrefix = []byte{0x19, 0x57}
		contractPrefix = []byte{0x19, 0x58}
	)

	if address == "" {
		return false
	}

	decodeBytes, err := utils.Base58Decode(address, base58Alphabet)
	if err != nil || len(decodeBytes) != 26 {
		return false
	}

	if (decodeBytes[0] != accountPrefix[0] || decodeBytes[1] != accountPrefix[1]) &&
		(decodeBytes[0] != contractPrefix[0] || decodeBytes[1] != contractPrefix[1]) {
		return false
	}

	check := owcrypt.Hash(decodeBytes[:22], 0, owcrypt.HASH_ALG_SHA3_256)[:4]

	for i := 0; i < 4; i ++ {
		if check[i] != decodeBytes[22 + i] {
			return false
		}
	}

	return true
}



