package numberofsubsequencesthatsatisfythegivensumcondition

import (
	"sort"
)

const mod int = 1000000007

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	pow2 := pow2Func()
	result := 0
	for start, num := range nums {
		endExclusive := start + sort.Search(len(nums)-start, func(i int) bool {
			return (nums[i+start] + num) > target
		})
		if start == endExclusive {
			continue
		}
		n := endExclusive - start - 1
		result += pow2(n)
		result %= mod
	}
	return result
}

// pow2Func returns a memoized function to calculate power of 2
func pow2Func() func(int) int {
	cache := map[int]int{0: 1, 1: 2}
	var pow2 func(int) int
	pow2 = func(n int) int {
		if cached, found := cache[n]; found {
			return cached
		}
		cache[n] = (2 * pow2(n-1)) % mod
		return cache[n]
	}
	return pow2
}
