package arithmeticslicesiisubsequence

func numberOfArithmeticSlices(nums []int) int {
	sliceLen := make([]map[int]int, len(nums))
	for i := range sliceLen {
		sliceLen[i] = make(map[int]int)
	}

	result := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			diff := nums[i] - nums[j]

			sliceLen[i][diff] += 1

			if lenAtJ, isPresent := sliceLen[j][diff]; isPresent {
				sliceLen[i][diff] += lenAtJ
				result += lenAtJ
			}
		}
	}
	return result
}
