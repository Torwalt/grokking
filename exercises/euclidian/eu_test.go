package euclidian_test

import (
	"grokking/exercises/euclidian"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindBiggestSquare(t *testing.T) {
	rec, err := euclidian.NewRectangle(1680, 640)
	require.NoError(t, err)
	require.NotEmpty(t, rec)

	biggest := euclidian.FindBiggestSquare(rec)
	require.NotEmpty(t, biggest)
	expRec := euclidian.Rectangle{Length: 80, Width: 80}
	assert.Equal(t, expRec, biggest)
}
