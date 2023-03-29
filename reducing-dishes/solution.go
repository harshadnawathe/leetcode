package reducingdishes

import (
	"sort"
)

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}

// Tabular DP
func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)

	maxScore := make([][]int, len(satisfaction)+1)
	for i := range maxScore {
		maxScore[i] = make([]int, len(satisfaction)+2)
	}

	for i := len(satisfaction) - 1; i >= 0; i-- {
		for time := len(satisfaction); time >= 1; time-- {
			maxScore[i][time] = maxOf(
				time*satisfaction[i]+maxScore[i+1][time+1],
				maxScore[i+1][time],
			)
		}
	}
	return maxScore[0][1]
}

// Recursion with memoisation
func maxSatisfaction2(satisfaction []int) int {
	sort.Ints(satisfaction)

	cache := make(map[key]int)

	var maxScore func(int, int) int
	maxScore = func(i, time int) int {
		if i == len(satisfaction) {
			return 0
		}
		if cached, found := cache[key{i, time}]; found {
			return cached
		}

		cache[key{i, time}] = maxOf(
			time*satisfaction[i]+maxScore(i+1, time+1),
			maxScore(i+1, time),
		)

		return cache[key{i, time}]
	}

	return maxScore(0, 1)
}

type key struct {
	i, time int
}

// Recursion
func maxSatisfaction1(satisfaction []int) int {
	sort.Ints(satisfaction)

	var maxScore func(int, int) int
	maxScore = func(i, time int) int {
		if i == len(satisfaction) {
			return 0
		}
		return maxOf(
			time*satisfaction[i]+maxScore(i+1, time+1),
			maxScore(i+1, time),
		)
	}

	return maxScore(0, 1)
}
