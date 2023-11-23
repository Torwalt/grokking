package selectionsort

type Artist struct {
	Name      string
	PlayCount int
}

func SelectionSort(artists []Artist) []Artist {
	out := make([]Artist, 0, len(artists))

	for len(artists) > 0 {
		idx := findSmallest(artists)
		out = append(out, artists[idx])

		// Assign last element to idx and reslice to remove last element.
		artists[idx] = artists[len(artists)-1]
		artists = artists[:len(artists)-1]
	}

	return out
}

func findSmallest(artists []Artist) int {
	idx := 0

	for i, a := range artists {
		if a.PlayCount <= artists[idx].PlayCount {
			idx = i
		}
	}

	return idx
}
