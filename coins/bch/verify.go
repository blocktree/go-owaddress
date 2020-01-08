package bch

import (
	"github.com/blocktree/go-owaddress/address"
	"github.com/blocktree/go-owaddress/utils"
	"github.com/blocktree/go-owcrypt"
	"strings"
)

// for register
var (
	DefaultStruct = &AddressVerify{}
	CoinName = "bch"
)

type AddressVerify struct {
	address.AddressVerify
}

func (b AddressVerify) IsValid (address string) bool {
	var(
		base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
		bech32Alphabet = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

		P2PKH_legacy_Prefix = byte(0x00)
		P2SH_legacy_Prefix = byte(0x05)
		Bech32_legacy_Prefix = "bitcoincash:q"

		P2PKH_copay_Prefix = byte(0x1C)
		P2SH_copay_Prefix = byte(0x28)
		Bech32_copay_Prefix = "bitcoincash:p"
	)

	if address == "" {
		return false
	}

	if strings.Index(address, Bech32_legacy_Prefix) == 0 || strings.Index(address, "q") == 0 { // p2pkh cash address

		if strings.Index(address, "q") == 0 {
			address = "bitcoincash:"+address
		}
		hash, err := Decode(address, bech32Alphabet)
		if err != nil || len(hash) != 21 || hash[0] != 0x00 {
			return false
		}

		return true
	}

	if strings.Index(address, Bech32_copay_Prefix) == 0 || strings.Index(address, "p") == 0 { // p2pkh cash address

		if strings.Index(address, "p") == 0 {
			address = "bitcoincash:"+address
		}
		hash, err := Decode(address, bech32Alphabet)
		if err != nil || len(hash) != 21 || hash[0] != 0x08 {
			return false
		}

		return true
	}

	if strings.Index(address, "1") == 0 {
		hash, err := utils.Base58Decode(address, base58Alphabet)
		if err != nil || hash[0] != P2PKH_legacy_Prefix {
			return false
		}
		check := owcrypt.Hash(hash[:21], 0, owcrypt.HASH_ALG_DOUBLE_SHA256)[:4]
		for key, value := range check {
			if value != hash[21+key] {
				return false
			}
		}

		return true
	}

	if strings.Index(address, "C") == 0 {
		hash, err := utils.Base58Decode(address, base58Alphabet)
		if err != nil || hash[0] != P2PKH_copay_Prefix {
			return false
		}

		check := owcrypt.Hash(hash[:21], 0, owcrypt.HASH_ALG_DOUBLE_SHA256)[:4]
		for key, value := range check {
			if value != hash[21+key] {
				return false
			}
		}
		return true
	}

	if strings.Index(address, "3") == 0 {
		hash, err := utils.Base58Decode(address, base58Alphabet)
		if err != nil || hash[0] != P2SH_legacy_Prefix {
			return false
		}

		check := owcrypt.Hash(hash[:21], 0, owcrypt.HASH_ALG_DOUBLE_SHA256)[:4]
		for key, value := range check {
			if value != hash[21+key] {
				return false
			}
		}
		return true
	}

	if strings.Index(address, "H") == 0 {
		hash, err := utils.Base58Decode(address, base58Alphabet)
		if err != nil || hash[0] != P2SH_copay_Prefix {
			return false
		}

		check := owcrypt.Hash(hash[:21], 0, owcrypt.HASH_ALG_DOUBLE_SHA256)[:4]
		for key, value := range check {
			if value != hash[21+key] {
				return false
			}
		}
		return true
	}

	return false
}

