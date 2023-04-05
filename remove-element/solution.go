package removeelement

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			right--
			for nums[right] == val && right > left {
				right--
			}

			nums[left], nums[right] = nums[right], nums[left]

		}
		left++
	}
	return right
}
