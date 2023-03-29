package sum

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	result := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		low, high := i+1, len(nums)-1
		for low < high {
			sum := nums[i] + nums[low] + nums[high]

			if sum < 0 {
				low++
			} else if sum > 0 {
				high--
			} else {
				result = append(result, []int{nums[i], nums[low], nums[high]})
				for low+1 < high && nums[low] == nums[low+1] {
					low++
				}
				for high-1 > low && nums[high] == nums[high-1] {
					high--
				}
				low++
				high--
			}
		}
	}
	return result
}
