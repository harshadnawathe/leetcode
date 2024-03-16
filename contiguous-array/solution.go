package contiguousarray

func findMaxLength(nums []int) int {
	diffIndex := make(map[int]int)
	diffIndex[0] = -1
	diff := 0

	maxLen := 0
	for i, num := range nums {
		if num == 0 {
			diff -= 1
		} else {
			diff += 1
		}

		if index, ok := diffIndex[diff]; ok {
			maxLen = maxOf(maxLen, i-index)
		} else {
			diffIndex[diff] = i
		}
	}
	return maxLen
}

func maxOf(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}
