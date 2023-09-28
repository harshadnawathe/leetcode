package sortarraybyparity

func sortArrayByParity(nums []int) []int {
	l, r := 0, len(nums)
	for l < r {
		if nums[l]&0x1 == 0 {
			l++
			continue
		}
		r--
		for l < r && nums[r]&0x1 == 1 {
			r--
		}

		nums[l], nums[r] = nums[r], nums[l]
	}

	return nums
}
