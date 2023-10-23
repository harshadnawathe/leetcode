package maximumscoreofagoodsubarray

func maximumScore(nums []int, k int) int {
	left, right := k, k
	min, maxScore := nums[k], nums[k]

	for left > 0 || right < len(nums)-1 {
		nextOnLeft := 0
		if left > 0 {
			nextOnLeft = nums[left-1]
		}
		nextOnRight := 0
		if right < len(nums)-1 {
			nextOnRight = nums[right+1]
		}

		if nextOnLeft < nextOnRight {
			right++
			min = minOf(min, nums[right])
		} else {
			left--
			min = minOf(min, nums[left])
		}

		maxScore = maxOf(maxScore, min*(right-left+1))
	}

	return maxScore
}

func minOf(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
