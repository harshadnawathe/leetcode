package minimumnumbersoffunctioncallstomaketargetarray

import "sort"

func minOperations(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	})

	numOperations := 0
	max := nums[0]
	if max == 0 {
		numOperations--
	}
	for max > 1 {
		if max&1 == 1 {
			max--
			numOperations++
		}
		max >>= 1
		numOperations++
	}

	for _, num := range nums[1:] {
		if num == 0 {
			numOperations--
		}
		for num > 1 {
			if num&1 == 1 {
				num--
				numOperations++
			}
			num >>= 1
		}
	}

	return numOperations + len(nums)
}
