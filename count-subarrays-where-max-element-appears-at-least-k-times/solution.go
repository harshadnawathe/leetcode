package countsubarrayswheremaxelementappearsatleastktimes

// time: O(n)
// space: O(1)
func countSubarrays(nums []int, k int) int64 {
	max := 0
	for _, num := range nums {
		if max < num {
			max = num
		}
	}

	subarrayCount := int64(0)
	begin := 0
	count := 0
	for end := range nums {
		if nums[end] == max {
			count++
		}

		for count == k {
			if nums[begin] == max {
				count--
			}
			begin++
		}

		subarrayCount += int64(begin)

	}

	return subarrayCount
}
