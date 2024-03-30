package commonsubstring

func CommonSubstring(in, other string) int {
	grid := make([][]int, len(in))
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, len(other))
	}

	out := 0
	for idxIn, letterIn := range in {
		for idxOther, letterOther := range other {
			if letterIn == letterOther {
				rowIdx := idxIn - 1
				colIDx := idxOther - 1
				if rowIdx < 0 || colIDx < 0 {
					// No neighbor.
					continue
				}

				v := grid[rowIdx][colIDx] + 1
				if v > out {
					out = v
				}
				grid[idxIn][idxOther] = v
			} else {
				grid[idxIn][idxOther] = 0
			}
		}
	}

	return out
}

func CommonSubsequence(in string, collection []string) []string {
	return []string{}
}
