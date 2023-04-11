package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	lb := search(len(nums), func(i int) bool {
		return nums[i] < target
	})
	if lb == len(nums) || nums[lb] != target {
		return []int{-1, -1}
	}

	ub := search(len(nums[lb:]), func(i int) bool {
		return nums[lb+i] <= target
	})

	if nums[lb+ub-1] != target {
		return []int{-1, -1}
	}

	return []int{lb, lb + ub - 1}
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
