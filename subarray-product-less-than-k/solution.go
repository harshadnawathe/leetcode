package subarrayproductlessthank

// time: O(n)
// space: O(1)
func numSubarrayProductLessThanK(nums []int, k int) int {
	count := 0

	for left, right, product := 0, 0, 1; right < len(nums); right++ {
		product *= nums[right]

		for left <= right && product >= k {
			product /= nums[left]
			left++
		}

		count += right - left + 1
	}

	return count
}
