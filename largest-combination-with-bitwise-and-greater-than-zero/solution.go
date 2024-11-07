package largestcombinationwithbitwiseandgreaterthanzero

func largestCombination(candidates []int) int {
	result := 0

	for i := range 24 {
		count := 0
		mask := 1 << i

		for _, candidate := range candidates {
			if mask&candidate != 0 {
				count++
			}
		}

		result = max(result, count)
	}

	return result
}
