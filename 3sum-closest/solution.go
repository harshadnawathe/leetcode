package sumclosest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	minDiff, result := math.MaxInt32, 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		low, high := i+1, len(nums)-1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]

			if sum < target {
				if diff := target - sum; diff < minDiff {
					minDiff = diff
					result = sum
				}
				low++
			} else if sum > target {
				if diff := sum - target; diff < minDiff {
					minDiff = diff
					result = sum
				}
				high--
			} else {
				low++
				high--

				minDiff = 0
				result = sum
			}
		}
	}
	return result
}
