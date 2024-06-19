package minimumnumberofdaystomakembouquets

import (
	"sort"
)

func minDays(bloomDay []int, m int, k int) int {
	n := 0
	for _, d := range bloomDay {
		n = max(n, d)
	}

	x, _ := sort.Find(n+1, func(d int) int {
		return m - numBouquets(bloomDay, k, d)
	})

	if x > n {
		return -1
	}

	return x
}

func numBouquets(bloomDay []int, k int, d int) int {
	num := 0

	for i, c := 0, 0; i < len(bloomDay); i++ {

		if bloomDay[i] <= d {
			c++
		} else {
			c = 0
		}

		if c == k {
			num++
			c = 0
		}
	}

	return num
}
