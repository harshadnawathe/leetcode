package findthepivotinteger

// time: O(n), space: O(1)
func pivotInteger(n int) int {
	if n == 1 {
		return 1
	}
	left, right := 1, n
	leftSum, rightSum := 1, n

	for left < right {
		if leftSum < rightSum {
			left++
			leftSum += left
		} else {
			right--
			rightSum += right
		}

		if leftSum == rightSum && left+1 == right-1 {
			return left + 1
		}
	}
	return -1
}
