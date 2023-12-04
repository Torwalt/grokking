package quicksort_test

import (
	"grokking/exercises/quicksort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		exp  []int
	}{
		{
			name: "test name",
			in:   []int{2, 5, 12, 1, 7, 9, 111, 4, 300000, 12},
			exp:  []int{1, 2, 4, 5, 7, 9, 12, 12, 111, 300000},
		},
		{
			name: "test name",
			in:   []int{2, 5, 12, 1, 7, 0, 9, 111, 4, 300000, 12},
			exp:  []int{0, 1, 2, 4, 5, 7, 9, 12, 12, 111, 300000},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := quicksort.QuickSort(test.in)
			require.Equal(t, test.exp, out)
		})
	}
}
