package minimumscorebychangingtwoelements

import "sort"

func minimizeSum(nums []int) int {
	sort.Ints(nums)

	return min3(
		abs(nums[2]-nums[len(nums)-1]),
		abs(nums[1]-nums[len(nums)-2]),
		abs(nums[0]-nums[len(nums)-3]),
	)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min3(x, y, z int) int {
	return min(x, min(y, z))
}

func min(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}
