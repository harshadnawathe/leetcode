package maximizehappinessofselectedchildren

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)

	sum := int64(0)

	for i, n := len(happiness)-1, 0; i >= 0 && k-n > 0; i, n = i-1, n+1 {
		if h := happiness[i] - n; h >= 0 {
			sum += int64(h)
		}
	}

	return sum
}
