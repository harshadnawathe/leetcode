package minimizemaximumofarray

func minimizeArrayValue(nums []int) int {
	prefixSum, result := 0, 0
	for i, num := range nums {
		prefixSum += num
		result = maxOf(result, (prefixSum+i)/(i+1))
	}
	return result
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
