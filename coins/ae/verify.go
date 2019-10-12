package ae

import (
	"github.com/blocktree/go-owaddress/address"
	"github.com/blocktree/go-owaddress/utils"
	"github.com/blocktree/go-owcrypt"
	"strings"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName = "ae"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {
	var(
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

		AEPrefixAccountPubkey = "ak_"
	)

	if address == "" {
		return false
	}

	var pubKeyMaterial string
	if strings.HasPrefix(address, AEPrefixAccountPubkey) {
		pubKeyMaterial = address[len(AEPrefixAccountPubkey):]
	}

	decodeBytes, err := utils.Base58Decode(pubKeyMaterial, base58Alphabet)
	if err != nil || len(decodeBytes) != 36 {
		return false
	}

	check := owcrypt.Hash(decodeBytes[:32], 0, owcrypt.HASh_ALG_DOUBLE_SHA256)[:4]

	for i := 0; i < 4; i ++ {
		if check[i] != decodeBytes[32 + i] {
			return false
		}
	}

	return true
}
