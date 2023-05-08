package largestrectangleinhistogram

func largestRectangleArea(heights []int) int {
	maxArea := 0
	leftLimit := leftLimits(heights)
	rightLimit := rightLimits(heights)
	for i, height := range heights {
		left, right := leftLimit[i], rightLimit[i]
		width := right - left + 1
		area := width * height
		if maxArea < area {
			maxArea = area
		}
	}
	return maxArea
}

// rightLimits
func rightLimits(heights []int) []int {
	limits := make([]int, len(heights))
	stk := stack{}
	for i := len(heights) - 1; i >= 0; i-- {
		height := heights[i]
		if len(stk) == 0 {
			limits[i] = i
		} else if heights[top(stk)] < height {
			limits[i] = top(stk) - 1
		} else {
			for {
				stk = pop(stk)

				if len(stk) == 0 {
					limits[i] = len(heights) - 1
					break
				}

				if heights[top(stk)] < height {
					limits[i] = top(stk) - 1
					break
				}
			}
		}
		stk = push(stk, i)
	}
	return limits
}

// leftLimits
func leftLimits(heights []int) []int {
	limits := make([]int, len(heights))
	stk := stack{}
	for i, height := range heights {
		if len(stk) == 0 {
			limits[i] = i
		} else if heights[top(stk)] < height {
			limits[i] = top(stk) + 1
		} else {
			for {
				stk = pop(stk)

				if len(stk) == 0 {
					limits[i] = 0
					break
				}

				if heights[top(stk)] < height {
					limits[i] = top(stk) + 1
					break
				}
			}
		}
		stk = push(stk, i)
	}
	return limits
}

// stack
type stack []int

func top(stk stack) int           { return stk[len(stk)-1] }
func pop(stk stack) stack         { return stk[:len(stk)-1] }
func push(stk stack, x int) stack { return append(stk, x) }
