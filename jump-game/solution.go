package jumpgame

// time: O(n)
// space: O(1)
func canJump(nums []int) bool {
	longest := 0

	for i, num := range nums {
		if longest < i {
			return false
		}

		longest = max(longest, i+num)
	}

	return true
}

// time: O(n^2)
// space: O(n)
func canJump1(nums []int) bool {
	isReachable := make([]bool, len(nums))
	isReachable[len(nums)-1] = true

	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] == 0 {
			isReachable[i] = false
			continue
		}

		for n := 1; n <= nums[i] && i+n < len(nums); n++ {
			if isReachable[i+n] {
				isReachable[i] = true
				break
			}
		}
	}

	return isReachable[0]
}
