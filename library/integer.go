package library

func NumberOfDigits(value int) int {
	dig := 0
	for value > 0 {
		dig++
		value = value / 10
	}
	return dig
}
