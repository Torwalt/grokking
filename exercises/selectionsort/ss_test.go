package selectionsort_test

import (
	"grokking/exercises/selectionsort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	artists := []selectionsort.Artist{
		{Name: "Megadeth", PlayCount: 666},
		{Name: "BabyMetal", PlayCount: 333},
		{Name: "Theocracy", PlayCount: 777},
		{Name: "Sabaton", PlayCount: 1337},
		{Name: "Dragonforce", PlayCount: 777},
		{Name: "Powerwolf", PlayCount: 999},
		{Name: "FFDP", PlayCount: 444},
		{Name: "Amon Amarth", PlayCount: 444},
		{Name: "Shitband", PlayCount: 1},
	}

	// Smallest first.
	expectedSorted := []selectionsort.Artist{
		{Name: "Shitband", PlayCount: 1},
		{Name: "BabyMetal", PlayCount: 333},
		{Name: "FFDP", PlayCount: 444},
		{Name: "Amon Amarth", PlayCount: 444},
		{Name: "Megadeth", PlayCount: 666},
		{Name: "Theocracy", PlayCount: 777},
		{Name: "Dragonforce", PlayCount: 777},
		{Name: "Powerwolf", PlayCount: 999},
		{Name: "Sabaton", PlayCount: 1337},
	}

	sortedArtists := selectionsort.SelectionSort(artists)

	assert.Equal(t, expectedSorted, sortedArtists)
}
