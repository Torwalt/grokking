package factorial_test

import (
	"grokking/exercises/factorial"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFactorize(t *testing.T) {
	tests := []struct {
		name string
		in   int
		exp  int
	}{
		{
			name: "test name",
			in:   5,
			exp:  120,
		},
		{
			name: "test name",
			in:   9,
			exp:  362880,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fac := factorial.Factorize(test.in)
			assert.Equal(t, test.exp, fac)
		})
	}
}
