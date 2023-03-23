package containerwithmostwater

func maxArea(height []int) int {
	result := 0
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			result = maxOf(result, height[left]*(right-left))
			left++
		} else {
			result = maxOf(result, height[right]*(right-left))
			right--
		}
	}
	return result
}

func maxOf(lhs, rhs int) int {
	if lhs < rhs {
		return rhs
	}
	return lhs
}
