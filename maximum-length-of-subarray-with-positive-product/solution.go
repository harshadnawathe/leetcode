package maximumlengthofsubarraywithpositiveproduct

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}

func getMaxLen(nums []int) int {
	var posLen, negLen, maxLen int
	for _, num := range nums {
		if num == 0 {
			posLen = 0
			negLen = 0
		} else {
			if num < 0 {
				posLen, negLen = negLen, posLen
			}
			if negLen > 0 || num < 0 {
				negLen++
			}
			if posLen > 0 || num > 0 {
				posLen++
			}
			maxLen = maxOf(maxLen, posLen)
		}
	}
	return maxLen
}

func getMaxLen2(nums []int) int {
	var maxLen func(int, int) int
	maxLen = func(begin, end int) int {
		if end == begin {
			return 0
		}

		product := 1
		for i := begin; i < end && product != 0; i++ {
			product *= nums[i]
		}
		if product > 0 {
			return end - begin
		}

		return maxOf(maxLen(begin+1, end), maxLen(begin, end-1))
	}

	return maxLen(0, len(nums))
}

func getMaxLen1(nums []int) int {
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			product := 1
			for k := i; k < j; k++ {
				product *= nums[k]
			}
			if product > 0 && (j-i) > maxLen {
				maxLen = j - i
			}
		}
	}
	return maxLen
}

/*                 3

             0 1 -2 -3 -4
               3/         \3
        0 1 -2 -3       1 -2 -3 -4
       1/     \3          3/     \2
   0 1 -2     1 -2 -3  1 -2 -3  -2 -3 -4
   /   \                          2/    \2
 0 1   1 -2                     -2 -3  -3 -4
 / \    /\
0   1  1  -2


     0  1 -2 -3 -4
  0  0  0 0
  1
 -2
 -3
 -4
*/
