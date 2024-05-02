package largestpositiveintegerthatexistswithitsnegative

func findMaxK(nums []int) int {
	complements := make(map[int]struct{})
	maxK := -1

	for _, num := range nums {
		if _, ok := complements[num]; ok {
			maxK = max(maxK, abs(num))
		} else {
			complements[-num] = struct{}{}
		}
	}

	return maxK
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
