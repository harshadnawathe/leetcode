package revealcardsinincreasingorder

import "sort"

// time: O(nlogn)
// space: O(n) - considering the space for result
func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)

	result := make([]int, len(deck))

	ir, id := 0, 0
	skip := false

	for id < len(deck) {
		if result[ir] == 0 {
			if !skip {
				result[ir] = deck[id]
				id++
			}
			skip = !skip
		}

		ir++
		if ir == len(result) {
			ir = 0
		}
	}

	return result
}
