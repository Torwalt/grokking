package commonsubstring_test

import (
	"grokking/exercises/commonsubstring"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommonSubstring(t *testing.T) {
	x := commonsubstring.CommonSubstring("hish", "fish")
	require.Equal(t, 3, x)
}
