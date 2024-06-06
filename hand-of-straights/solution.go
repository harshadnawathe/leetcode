package handofstraights

import "sort"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	sort.Ints(hand)

	numCount := make(map[int]int)
	for _, num := range hand {
		numCount[num]++
	}

	numGroups := 0

	for _, num := range hand {
		if numCount[num] == 0 {
			continue
		}

		numCount[num]--
		for i := 1; i < groupSize; i++ {
			if numCount[num+i] == 0 {
				return false
			}

			numCount[num+i]--
		}

		numGroups++
	}

	return numGroups == len(hand)/groupSize
}
