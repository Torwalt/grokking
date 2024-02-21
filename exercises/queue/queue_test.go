package queue_test

import (
	"grokking/exercises/queue"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	q, err := queue.NewQueue(20)
	require.NoError(t, err)

	q.Enqueue(4)
}
