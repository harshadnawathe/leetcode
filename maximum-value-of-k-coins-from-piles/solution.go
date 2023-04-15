package maximumvalueofkcoinsfrompiles

func maxValueOfCoins(piles [][]int, k int) int {
	maxValue := make([][]int, len(piles)+1)
	for i := range maxValue {
		maxValue[i] = make([]int, k+1)
	}

	for i := 1; i <= len(piles); i++ {
		for coins := 1; coins <= k; coins++ {
			maxValue[i][coins] = maxValue[i-1][coins]

			sum := 0
			for c := 0; c < minOf(len(piles[i-1]), coins); c++ {
				sum += piles[i-1][c]
				maxValue[i][coins] = maxOf(maxValue[i][coins], sum+maxValue[i-1][coins-(c+1)])
			}
		}
	}

	return maxValue[len(piles)][k]
}

// simple recursion
func maxValueOfCoins1(piles [][]int, k int) int {
	var maxValue func(int, int) int
	maxValue = func(i, coins int) int {
		if i == 0 || coins == 0 {
			return 0
		}

		// select no coins from this pile
		max := maxValue(i-1, coins)
		// select coins one by one and recurse
		sum := 0
		for c := 0; c < minOf(len(piles[i-1]), coins); c++ {
			sum += piles[i-1][c]
			max = maxOf(max, sum+maxValue(i-1, coins-(c+1)))
		}
		return max
	}

	return maxValue(len(piles), k)
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
