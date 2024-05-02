package minimumoperationstoformsubsequencewithtargetsum

import "sort"

func minOperations(nums []int, target int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	if target > total {
		return -1
	}

	sort.Ints(nums)

	count := 0

	for target > 0 {
		num := nums[len(nums)-1]
		nums = nums[:len(nums)-1]

		if total-num >= target {
			total -= num
		} else if num <= target {
			total -= num
			target -= num
		} else {
			nums = append(nums, num/2, num/2)
			count++
		}
	}

	return count
}
