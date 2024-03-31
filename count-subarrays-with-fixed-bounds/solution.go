package countsubarrayswithfixedbounds

func countSubarrays(nums []int, minK int, maxK int) int64 {
	count := int64(0)
	left, right, bad := -1, -1, -1

	for i, num := range nums {
		if num < minK || num > maxK {
			bad = i
		}

		if num == minK {
			left = i
		}

		if num == maxK {
			right = i
		}

		count += int64(maxOf(0, minOf(left, right)-bad))
	}
	return count
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}

func minOf(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}
