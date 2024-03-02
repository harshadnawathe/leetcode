package squaresofasortedarray

// time: O(n), space: O(n)
func sortedSquares(nums []int) []int {
	indexNonNegative := len(nums)
	for i, num := range nums {
		if num > -1 {
			indexNonNegative = i
			break
		}
	}

	left, right := indexNonNegative-1, indexNonNegative
	result := make([]int, len(nums))
	for i := range result {
		if left >= 0 && right < len(nums) {
			leftSquare, rightSquare := nums[left]*nums[left], nums[right]*nums[right]
			if leftSquare < rightSquare {
				result[i] = leftSquare
				left--
			} else {
				result[i] = rightSquare
				right++
			}
		} else if left >= 0 {
			result[i] = nums[left] * nums[left]
			left--
		} else if right < len(nums) {
			result[i] = nums[right] * nums[right]
			right++
		}
	}

	return result
}
