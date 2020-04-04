package ark

// for register
var (
	DefaultStruct = &Verifier{}
	CoinName      = "ark"
)

type Verifier struct{}

func (b Verifier) IsValid(address string) bool {

	return true
}
