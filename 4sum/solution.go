package sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	return kSum(4, nums, target)
}

func kSum(k int, nums []int, target int) [][]int {
	sort.Ints(nums)
	return kSumSorted(k, nums, target)
}

func kSumSorted(k int, nums []int, target int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	if avg := target / k; nums[0] > avg || nums[len(nums)-1] < avg {
		return nil
	}

	if k == 2 {
		return twoSumSorted(nums, target)
	}

	var result [][]int
	for i, num := range nums {
		if i > 0 && nums[i-1] == num {
			continue
		}

		for _, r := range kSumSorted(k-1, nums[i+1:], target-num) {
			result = append(result, append([]int{num}, r...))
		}

	}
	return result
}

func twoSumSorted(nums []int, target int) [][]int {
	pairs := [][]int{}

	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]

		if sum < target || (left-1 >= 0 && nums[left] == nums[left-1]) {
			left++
		} else if sum > target || (right+1 < len(nums) && nums[right] == nums[right+1]) {
			right--
		} else {
			pairs = append(pairs, []int{nums[left], nums[right]})
			left++
			right--
		}
	}

	return pairs
}
