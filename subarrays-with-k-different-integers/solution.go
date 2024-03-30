package subarrayswithkdifferentintegers

// time: O(n)
// space: O(n)
func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithAtMostKDistinct(nums, k) - subarraysWithAtMostKDistinct(nums, k-1)
}

// time: O(n)
// space: O(n)
func subarraysWithAtMostKDistinct(nums []int, k int) int {
	count := 0

	numToCount := make(map[int]int)
	for begin, end := 0, 0; end < len(nums); end++ {
		numToCount[nums[end]] += 1

		for len(numToCount) > k {
			if numToCount[nums[begin]] == 1 {
				delete(numToCount, nums[begin])
			} else {
				numToCount[nums[begin]] -= 1
			}
			begin++
		}

		count += end - begin + 1
	}

	return count
}
