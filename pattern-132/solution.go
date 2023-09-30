package pattern132

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	min := make([]int, len(nums))
	min[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		min[i] = minOf(min[i-1], nums[i])
	}
	for j, k := len(nums)-1, len(nums); j >= 0; j-- {
		if nums[j] > min[j] {
			for k < len(nums) && nums[k] <= min[j] {
				k++
			}
			if k < len(nums) && nums[k] < nums[j] {
				return true
			}
			k--
			nums[k] = nums[j]
		}
	}
	return false
}

func minOf(a, b int) int {
	if b < a {
		return b
	}
	return a
}
