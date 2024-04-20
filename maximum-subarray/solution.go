package maximumsubarray

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	sumSoFar := max(0, nums[0])

	for _, num := range nums[1:] {
		sumSoFar += num

		maxSum = max(maxSum, sumSoFar)

		sumSoFar = max(0, sumSoFar)
	}

	return maxSum
}
