package lengthoflongestsubarraywithatmostkfrequency

// time: O(n)
// space: O(n)
func maxSubarrayLength(nums []int, k int) int {
	maxLen := 0

	numCount := make(map[int]int)

	for left, right := 0, 0; right < len(nums); right++ {
		numCount[nums[right]]++

		for numCount[nums[right]] > k {
			numCount[nums[left]]--
			left++
		}

		if currLen := right - left + 1; currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}
