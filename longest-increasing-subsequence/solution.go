package longestincreasingsubsequence

import (
	"sort"
)

func lengthOfLIS(nums []int) int {
	maxLen := 0
	for _, num := range nums {
		n := sort.Search(maxLen, func(i int) bool {
			return nums[i] >= num
		})

		if n == maxLen {
			maxLen++
		}

		nums[n] = num
	}
	return maxLen
}
