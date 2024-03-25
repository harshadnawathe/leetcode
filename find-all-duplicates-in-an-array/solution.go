package findallduplicatesinanarray

func findDuplicates(nums []int) []int {
	dups := []int{}

	for _, num := range nums {
		if num < 0 {
			num = -num
		}

		if nums[num-1] < 0 {
			dups = append(dups, num)
		} else {
			nums[num-1] *= -1
		}
	}

	return dups
}
