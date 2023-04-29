package findpeakelement

func findPeakElement(nums []int) int {
	count := len(nums) - 1
	first := 0
	for count > 0 {
		step := count / 2
		mid := first + step
		if nums[mid] > nums[mid+1] {
			count = step
		} else {
			first = mid + 1
			count -= (step + 1)
		}
	}
	return first
}

// Time: O(log(n))
// Space: O(log(n)) // stack
func findPeakElement2(nums []int) int {
	var find func(int, int) int
	find = func(first, last int) int {
		if first == last {
			return first
		}

		mid := (last + first) / 2
		if nums[mid] > nums[mid+1] {
			return find(first, mid)
		}
		return find(mid+1, last)
	}
	return find(0, len(nums)-1)
}

// Time: O(n)
// Space: O(1)
func findPeakElement1(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}
