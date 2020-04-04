package owaddress

import (
	"errors"
	"strings"

	"github.com/star001007/go-owaddress/address"
)

func getVerifier(coin string) (address.AddressVerifier, error) {
	verifier, ok := AddressVerifierRegistry[coin]
	if !ok {
		return nil, errors.New("Coin [" + coin + "] is not registered!")
	}

	return verifier, nil
}

func Verify(coin, address string) (bool, error) {
	verifier, err := getVerifier(strings.ToLower(coin))
	if err != nil {
		return false, err
	}
	return verifier.IsValid(address), nil
}
