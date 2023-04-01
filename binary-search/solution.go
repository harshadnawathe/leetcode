package binarysearch

func search(nums []int, target int) int {
	first, count := 0, len(nums)

	for count > 0 {
		step := count / 2
		mid := first + step

		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			count = step
		} else {
			first = mid + 1
			count -= (step + 1)
		}
	}
	return -1
}
