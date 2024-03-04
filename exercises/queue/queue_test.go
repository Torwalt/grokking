package queue_test

import (
	"grokking/exercises/queue"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	g := queue.Graph{
		"Alex":   []string{"Annika", "Georg", "Irina"},
		"Annika": []string{"Astrid", "Guido", "Nadine", "Alex"},
		"Nadine": []string{"Astrid", "Guido", "Annika", "Arthur"},
		"Georg":  []string{"Alex", "Sasha"},
		"Irina":  []string{"Alex", "Leyla"},
		"Astrid": []string{"Guido", "Annika", "Nadine"},
	}

	found := queue.BreadthFirstSearch(g, "Alex", "Leyla")
	require.True(t, found)
}
