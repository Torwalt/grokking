package binarysearch_test

import (
	"grokking/exercises/binarysearch"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		size   int
		search int
	}{
		{
			name:   "test name",
			size:   100,
			search: 34,
		},
		{
			name:   "test name",
			size:   17293,
			search: 4793,
		},
		{
			name:   "test name",
			size:   123131,
			search: 6789,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			in := make([]int, 0, test.size)
			for i := 0; i < test.size; i++ {
				in = append(in, i)
			}

			idx, found := binarysearch.BinarySearch(in, test.search)

			require.True(t, found)
			assert.Equal(t, idx, test.search)
			assert.Equal(t, test.search, in[idx])
		})
	}
}
