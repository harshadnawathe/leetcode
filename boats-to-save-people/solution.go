package boatstosavepeople

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	left, right := 0, len(people)-1
	boatCount := 0
	for left <= right {
		boatCount++

		if totalWeight := people[left] + people[right]; totalWeight <= limit {
			left++
		}
		right--
	}
	return boatCount
}
