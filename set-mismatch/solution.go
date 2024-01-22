package setmismatch

// time O(n) | space O(1)
func findErrorNums(nums []int) []int {
	result := []int{}
	for i := range nums {
		num := abs(nums[i])
		if nums[num-1] < 0 {
			result = append(result, num)
		} else {
			nums[num-1] *= -1
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			result = append(result, i+1)
			break
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
