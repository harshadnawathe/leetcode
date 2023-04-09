package nextpermutation

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	reverse(nums[i+1:])
}

func reverse(nums []int) {
	left, right := 0, len(nums)
	for left < right {
		right--
		nums[left], nums[right] = nums[right], nums[left]
		left++
	}
}
