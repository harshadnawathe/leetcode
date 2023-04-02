package successfulpairsofspellsandpotions

import (
	"sort"
)

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)

	pairsCount := make([]int, len(spells))
	for i, spell := range spells {

		minSuccess := sort.Search(len(potions), func(i int) bool {
			return int64(potions[i])*int64(spell) >= success
		})

		pairsCount[i] = len(potions) - minSuccess
	}
	return pairsCount
}

// Brute force
func successfulPairs1(spells []int, potions []int, success int64) []int {
	pairsCount := make([]int, len(spells))
	for i, spell := range spells {
		for _, potion := range potions {
			if int64(spell)*int64(potion) >= success {
				pairsCount[i]++
			}
		}
	}
	return pairsCount
}
