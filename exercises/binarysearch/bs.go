package binarysearch

func BinarySearch(in []int, x int) (int, bool) {
	// First index.
	low := 0
	// Highest index.
	high := len(in) - 1

	for low <= high {
		mid := (high + low) / 2
		current := in[mid]

		if current == x {
			return mid, true
		}

		if current > x {
			high = current
			continue
		}

		low = current
	}

	return 0, false
}
