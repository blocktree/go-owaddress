package owaddress

import (
	"errors"
	"reflect"
)

func getObject(coin string) (interface{}, error) {
	elem, ok  := AddressVerifyRegistry[coin]
	if !ok {
		return nil, errors.New("Coin [" + coin + "] is not registered!")
	}

	return reflect.New(elem).Elem().Interface(), nil
}

func Verify(coin, address string) (bool, error) {

	object, err := getObject(coin)
	if err != nil {
		return false, err
	}

	input := reflect.ValueOf(&address)
	input = reflect.Indirect(input)
	input.SetString(address)

	ret := reflect.ValueOf(object).MethodByName("IsValid").Call([]reflect.Value{input})

	return ret[0].Bool(), nil
}
