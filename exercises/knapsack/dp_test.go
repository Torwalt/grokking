package knapsack_test

import (
	"grokking/exercises/knapsack"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptimizeLoot(t *testing.T) {
	tests := []struct {
		name             string
		inLoot           []knapsack.Loot
		expOptimizedLoot []knapsack.Loot
		weight           uint
	}{
		{
			name: "normal case",
			inLoot: []knapsack.Loot{
				{
					Name:   "Guitar",
					Price:  1500,
					Weight: 1,
				},
				{
					Name:   "Laptop",
					Price:  2000,
					Weight: 3,
				},
				{
					Name:   "Stereo",
					Price:  3000,
					Weight: 4,
				},
			},
			expOptimizedLoot: []knapsack.Loot{
				{
					Name:   "Guitar",
					Price:  1500,
					Weight: 1,
				},
				{
					Name:   "Laptop",
					Price:  2000,
					Weight: 3,
				},
			},
			weight: 4,
		},
		{
			name: "reversing the loot should not change result",
			inLoot: []knapsack.Loot{
				{
					Name:   "Stereo",
					Price:  3000,
					Weight: 4,
				},
				{
					Name:   "Laptop",
					Price:  2000,
					Weight: 3,
				},
				{
					Name:   "Guitar",
					Price:  1500,
					Weight: 1,
				},
			},
			expOptimizedLoot: []knapsack.Loot{
				{
					Name:   "Guitar",
					Price:  1500,
					Weight: 1,
				},
				{
					Name:   "Laptop",
					Price:  2000,
					Weight: 3,
				},
			},
			weight: 4,
		},
		{
			name: "add iphone",
			inLoot: []knapsack.Loot{
				{
					Name:   "Stereo",
					Price:  3000,
					Weight: 4,
				},
				{
					Name:   "Laptop",
					Price:  2000,
					Weight: 3,
				},
				{
					Name:   "iphone",
					Price:  2000,
					Weight: 1,
				},
				{
					Name:   "Guitar",
					Price:  1500,
					Weight: 1,
				},
			},
			expOptimizedLoot: []knapsack.Loot{
				{
					Name:   "Laptop",
					Price:  2000,
					Weight: 3,
				},
				{
					Name:   "iphone",
					Price:  2000,
					Weight: 1,
				},
			},
			weight: 4,
		},
		{
			name: "add mp3 player",
			inLoot: []knapsack.Loot{
				{
					Name:   "Stereo",
					Price:  3000,
					Weight: 4,
				},
				{
					Name:   "MP3 Player",
					Price:  1000,
					Weight: 1,
				},
				{
					Name:   "Laptop",
					Price:  2000,
					Weight: 3,
				},
				{
					Name:   "Guitar",
					Price:  1500,
					Weight: 1,
				},
				{
					Name:   "iphone",
					Price:  2000,
					Weight: 1,
				},
			},
			expOptimizedLoot: []knapsack.Loot{
				{
					Name:   "MP3 Player",
					Price:  1000,
					Weight: 1,
				},
				{
					Name:   "Guitar",
					Price:  1500,
					Weight: 1,
				},
				{
					Name:   "iphone",
					Price:  2000,
					Weight: 1,
				},
			},
			weight: 4,
		},
		{
			name: "different example",
			inLoot: []knapsack.Loot{
				{
					Name:   "Water",
					Price:  10,
					Weight: 3,
				},
				{
					Name:   "Book",
					Price:  3,
					Weight: 1,
				},
				{
					Name:   "Food",
					Price:  9,
					Weight: 2,
				},
				{
					Name:   "Jacket",
					Price:  5,
					Weight: 2,
				},
				{
					Name:   "Camera",
					Price:  6,
					Weight: 1,
				},
			},
			expOptimizedLoot: []knapsack.Loot{
				{
					Name:   "Water",
					Price:  10,
					Weight: 3,
				},
				{
					Name:   "Food",
					Price:  9,
					Weight: 2,
				},
				{
					Name:   "Camera",
					Price:  6,
					Weight: 1,
				},
			},
			weight: 6,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			optimizedLoot := knapsack.OptimizeLoot(test.inLoot, test.weight)
			assert.Equal(t, knapsack.SortLoot(test.expOptimizedLoot), knapsack.SortLoot(optimizedLoot))
		})
	}
}
