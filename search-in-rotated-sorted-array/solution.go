package searchinrotatedsortedarray

func search(nums []int, target int) int {
	first, count := 0, len(nums)
	for count > 0 {
		step := count / 2
		mid := first + step

		if nums[mid] == target {
			return mid
		}

		if nums[first] < nums[mid] {
			if target < nums[mid] && nums[first] <= target {
				count = step
			} else {
				first = mid + 1
				count -= (step + 1)
			}
		} else {
			if nums[mid] < target && target <= nums[first+count-1] {
				first = mid + 1
				count -= (step + 1)
			} else {
				count = step
			}
		}
	}
	return -1
}
