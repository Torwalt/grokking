package hashmap_test

import (
	"grokking/exercises/hashmap"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	m := hashmap.NewMap(13)
	require.NotEmpty(t, m)

	m.Insert("apple", 10)
	m.Insert("ananas", 17)
	m.Insert("Cupra Formentor VZ5 Taiga Grey", 58000)
	assert.Equal(t, 3, m.Count())

	// For the same input, there must be the same index, thus the count should
	// not have changed.
	m.Insert("ananas", 7)
	assert.Equal(t, 3, m.Count())

	assertPrice := func(key string, expPrice int) {
		price, ok := m.Get(key)
		assert.True(t, ok)
		assert.Equal(t, expPrice, price)
	}

	assertPrice("apple", 10)
	assertPrice("ananas", 7)
	assertPrice("Cupra Formentor VZ5 Taiga Grey", 58000)

	_, exists := m.Get("Deathstar")
	assert.False(t, exists)
}
