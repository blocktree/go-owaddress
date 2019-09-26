package address

type AddressVerify interface {
	IsValid(coin, address string) bool
}
