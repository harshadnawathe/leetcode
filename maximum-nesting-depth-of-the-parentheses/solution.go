package maximumnestingdepthoftheparentheses

func maxDepth(s string) int {
	maxDepth := 0
	currDepth := 0

	for _, char := range s {
		if char == '(' {
			currDepth++
			maxDepth = max(maxDepth, currDepth)
		} else if char == ')' {
			currDepth--
		}
	}

	return maxDepth
}
