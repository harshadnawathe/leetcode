package trappingrainwater

// time: O(n), space: O(1)
func trap(height []int) int {
	totalWater := 0

	leftMax, rightMax := 0, 0
	left, right := 0, len(height)-1
	for left < right {
		if lh, rh := height[left], height[right]; lh < rh {
			if lh >= leftMax {
				leftMax = lh
			} else {
				totalWater += leftMax - lh
			}
			left++
		} else {
			if rh >= rightMax {
				rightMax = rh
			} else {
				totalWater += rightMax - rh
			}
			right--
		}
	}

	return totalWater
}

// time: O(n), space: O(n)
func trap1(height []int) int {
	leftMax := make([]int, len(height))
	for i := 1; i < len(height); i++ {
		leftMax[i] = maxOf(leftMax[i-1], height[i-1])
	}

	rightMax := make([]int, len(height))
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = maxOf(rightMax[i+1], height[i+1])
	}

	totalWater := 0
	for i, h := range height {
		if water := minOf(leftMax[i], rightMax[i]) - h; water > 0 {
			totalWater += water
		}
	}

	return totalWater
}

func minOf(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
