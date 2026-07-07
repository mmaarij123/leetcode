import "strconv"

func sumAndMultiply(n int) int64 {
	x := int64(0)
	sum := int64(0)

	for _, ch := range strconv.Itoa(n) {
		digit := int64(ch - '0')

		if digit != 0 {
			x = x*10 + digit
			sum += digit
		}
	}

	return x * sum
}