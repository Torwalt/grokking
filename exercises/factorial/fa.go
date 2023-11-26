package factorial

func Factorize(x int) int {
	if x == 1 {
		return 1
	}

	return x * Factorize(x-1)
}
