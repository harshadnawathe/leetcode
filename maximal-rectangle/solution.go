package maximalrectangle

// time: O(r * c)
// space: O(r * c)
func maximalRectangle(matrix [][]byte) int {
	hists := make([]histogram, len(matrix))
	for i := range matrix {
		hists[i] = make(histogram, len(matrix[i]))
	}

	// populate first row
	for i, b := range matrix[0] {
		hists[0][i] = int(b - '0')
	}

	// populate remaining rows
	for r := 1; r < len(matrix); r++ {
		for c, b := range matrix[r] {
			if b == '0' {
				hists[r][c] = 0
			} else {
				hists[r][c] = hists[r-1][c] + 1
			}
		}
	}

	maxArea := 0
	for _, hist := range hists {
		maxArea = max(maxArea, maxRectArea(hist))
	}

	return maxArea
}

// histogram
type histogram []int

func maxRectArea(hist histogram) int {
	leftLimit := leftLimit(hist)

	rightLimit := rightLimit(hist)

	maxArea := 0

	for i, h := range hist {
		area := h * (1 + rightLimit[i] - leftLimit[i])
		maxArea = max(maxArea, area)
	}

	return maxArea
}

func leftLimit(hist histogram) []int {
	limit := make([]int, len(hist))
	stk := stack{}

	for i := len(hist) - 1; i >= 0; i-- {
		h := hist[i]

		t, ok := top(stk)
		for ok && h < hist[t] {
			limit[t] = i + 1

			stk = pop(stk)
			t, ok = top(stk)
		}

		stk = push(stk, i)
	}

	return limit
}

func rightLimit(hist histogram) []int {
	limit := make([]int, len(hist))
	stk := stack{}

	for i, h := range hist {
		t, ok := top(stk)
		for ok && h < hist[t] {
			limit[t] = i - 1

			stk = pop(stk)
			t, ok = top(stk)
		}

		stk = push(stk, i)
	}

	for _, i := range stk {
		limit[i] = len(hist) - 1
	}

	return limit
}

// stack
type stack []int

func push(stk stack, val int) stack { return append(stk, val) }
func pop(stk stack) stack           { return stk[:len(stk)-1] }
func top(stk stack) (int, bool) {
	if len(stk) == 0 {
		return 0, false
	}
	return stk[len(stk)-1], true
}
