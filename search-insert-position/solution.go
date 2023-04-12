package searchinsertposition

func searchInsert(nums []int, target int) int {
	return search(len(nums), func(i int) bool {
		return nums[i] < target
	})
}

func search(n int, f func(int) bool) int {
	first, count := 0, n

	for count > 0 {
		step := count / 2
		mid := first + step

		if !f(mid) {
			count = step
		} else {
			count -= (step + 1)
			first = mid + 1
		}
	}
	return first
}
