package findtriangularsumofanarray

// time: O(n^2)
// space: O(1)
func triangularSum(nums []int) int {
	for len(nums) > 1 {
		newLen := len(nums) - 1
		for j := 0; j < newLen; j++ {
			nums[j] = nums[j] + nums[j+1]
			nums[j] %= 10
		}
		nums = nums[:newLen]
	}
	return nums[0]
}
