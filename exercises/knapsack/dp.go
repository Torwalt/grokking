package knapsack

import (
	"fmt"
	"sort"
)

// Solving a knapsack problem with dynamic programming (DP).

// We are a thief and we want to steal the most expensive items. But we can only carry so much.

type Kilogram = uint

type Loot struct {
	Name   string
	Price  uint
	Weight Kilogram
}

type AggregatedLoot struct {
	Value   uint
	Weight  uint
	NameIdx map[string]struct{}
}

// Sorts inplace
func SortLoot(in []Loot) []Loot {
	sort.Slice(in, func(a, b int) bool {
		return in[a].Name < in[b].Name
	})
	return in
}

func LootCheck(in []Loot) AggregatedLoot {
	out := AggregatedLoot{
		NameIdx: map[string]struct{}{},
	}

	for _, l := range in {
		out.Value += l.Price
		out.Weight += l.Weight
		out.NameIdx[l.Name] = struct{}{}
	}
	return out
}

func (l Loot) String() string {
	return fmt.Sprintf("|%v %v€ %vKG|", l.Name, l.Price, l.Weight)
}

func OptimizeLoot(toSteal []Loot, maxWeight uint) []Loot {
	lootIdx := map[string]struct{}{}
	uniqueLoot := []Loot{}
	for _, l := range toSteal {
		if _, ok := lootIdx[l.Name]; ok {
			continue
		}
		uniqueLoot = append(uniqueLoot, l)
	}

	wheightIdx := map[Kilogram][]Loot{}
	printWeightIdx := func(wi map[Kilogram][]Loot) {
		for kg, loots := range wi {
			fmt.Printf("KG: %v, loot: %v\n", kg, loots)
		}
	}

	// Sort uniqueLoot so that objectively better Loot comes before worse loot.
	sort.Slice(uniqueLoot, func(i, j int) bool {
		return uniqueLoot[i].Weight <= uniqueLoot[j].Weight
	})
	sort.Slice(uniqueLoot, func(i, j int) bool {
		return uniqueLoot[i].Price >= uniqueLoot[j].Price
	})

	for _, l := range uniqueLoot {
		for i := uint(1); i <= maxWeight; i++ {
			if l.Weight > i {
				continue
			}

			if _, ok := wheightIdx[i]; !ok {
				wheightIdx[i] = []Loot{l}
				continue
			}

			agg := LootCheck(wheightIdx[i])

			// Do not add the same Item again.
			if _, ok := agg.NameIdx[l.Name]; ok {
				continue
			}

			isEnoughSpace := (agg.Weight + l.Weight) <= i
			if isEnoughSpace {
				wheightIdx[i] = append(wheightIdx[i], l)
				continue
			}

			// Need to compare Price with potential Price of the next highest Loot slot.
			potentialLeftSpace := i - l.Weight
			otherLoot := wheightIdx[potentialLeftSpace]
			otherAgg := LootCheck(otherLoot)

			// Do not add the same Item again.
			if _, ok := otherAgg.NameIdx[l.Name]; ok {
				continue
			}

			if agg.Value >= l.Price+otherAgg.Value {
				continue
			}

			wheightIdx[i] = append([]Loot{l}, otherLoot...)
		}
	}

	printWeightIdx(wheightIdx)

	// Retrieve the actual biggest loot, as the one with the maxWeight is not always the best.
	biggestAgg := LootCheck([]Loot{})
	biggestLoot := []Loot{}
	for _, loots := range wheightIdx {
		agg := LootCheck(loots)
		if agg.Value > biggestAgg.Value {
			biggestLoot = loots
			biggestAgg = agg
		}
	}

	return biggestLoot
}
