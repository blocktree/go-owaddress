package address

type AddressVerifier interface {
	IsValid(address string) bool
}
