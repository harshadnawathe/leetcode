package removeduplicatesfromsortedarrayii

func removeDuplicates(nums []int) int {
	if n := len(nums); n <= 2 {
		return n
	}

	left := 2
	for right := 2; right < len(nums); right++ {
		if x := nums[right]; x != nums[left-1] || x != nums[left-2] {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}
