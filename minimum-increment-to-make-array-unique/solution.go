package minimumincrementtomakearrayunique

import "sort"

func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)

	var totalInc int

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			inc := nums[i-1] + 1 - nums[i]

			totalInc += inc
			nums[i] = nums[i-1] + 1
		}
	}

	return totalInc
}
