package utils

func IsLowerHex(data string) bool {

	if data == "" || len(data) % 2 == 1 {
		return false
	}

	for _, d := range data {
		if !((d >= '0' && d <= '9') || (d >= 'a' && d <= 'f')) {
			return false
		}
	}

	return true
}
