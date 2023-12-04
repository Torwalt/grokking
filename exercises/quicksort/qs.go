package quicksort

func QuickSort(in []int) []int {
	if len(in) <= 1 {
		return in
	}

	if len(in) == 2 {
		if in[0] > in[0] {
			in[0], in[1] = in[1], in[0]
		}

		return in
	}

	half := len(in) / 2
	pivotIdx := half - 1
	pivotElem := in[pivotIdx]

	left := make([]int, 0, len(in)-half)
	right := make([]int, 0, len(in)-half)

	for idx, e := range in {
		if idx == pivotIdx {
			continue
		}

		if e <= pivotElem {
			left = append(left, e)
			continue
		}

		right = append(right, e)
	}

	out := make([]int, 0, len(in))
	out = append(out, QuickSort(left)...)
	out = append(out, pivotElem)
	out = append(out, QuickSort(right)...)

	return out
}
