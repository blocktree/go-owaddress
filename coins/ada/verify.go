package ada

const addressLength = 20

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "cardano"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {
	if address == "" {
		return false
	}
	return true
}
