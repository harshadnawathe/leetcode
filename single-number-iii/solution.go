package singlenumberiii

func singleNumber(nums []int) []int {
	acc := 0
	for _, num := range nums {
		acc ^= num
	}

	acc &= -acc

	result := []int{0, 0}

	for _, num := range nums {
		if acc&num == 0 {
			result[0] ^= num
		} else {
			result[1] ^= num
		}
	}

	return result
}
