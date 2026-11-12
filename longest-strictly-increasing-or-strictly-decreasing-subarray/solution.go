package longeststrictlyincreasingorstrictlydecreasingsubarray

func longestMonotonicSubarray(nums []int) int {
	incLen, decLen, maxLen := 1, 1, 1

	for i := 1; i < len(nums); i++ {
		if diff := nums[i] - nums[i-1]; diff > 0 {
			incLen += 1
			decLen = 1
		} else if diff < 0 {
			incLen = 1
			decLen += 1
		} else {
			incLen, decLen = 1, 1
		}

		if currMax := max(incLen, decLen); maxLen < currMax {
			maxLen = currMax
		}
	}

	return maxLen
}
