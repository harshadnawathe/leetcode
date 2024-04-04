package rotatearray

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	}
	k = len(nums) - k
	begin, end := 0, len(nums)
	mid := k
	for begin != mid {
		nums[begin], nums[mid] = nums[mid], nums[begin]
		begin++
		mid++

		if mid == end {
			mid = k
		} else if begin == k {
			k = mid
		}
	}
}
