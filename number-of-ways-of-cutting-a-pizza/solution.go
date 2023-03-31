package numberofwaysofcuttingapizza

const mod = 1000000007

func applesCount(pizza []string) [][]int {

	apples := make([][]int, len(pizza)+1)
	for i := range apples {
		apples[i] = make([]int, len(pizza[0])+1)
	}

	for row := len(pizza) - 1; row >= 0; row-- {
		for col := len(pizza[0]) - 1; col >= 0; col-- {
			apples[row][col] = apples[row+1][col] + apples[row][col+1] - apples[row+1][col+1]
			if pizza[row][col] == 'A' {
				apples[row][col] += 1
			}
		}
	}
	return apples
}

func ways(pizza []string, k int) int {
	apples := applesCount(pizza)

	cache := make(map[[3]int]int)

	var numWays func(top, left int, k int) int
	numWays = func(top, left int, k int) int {
		key := [3]int{top, left, k}
		if cached, found := cache[key]; found {
			return cached
		}

		if k == 1 {
			if apples[top][left] > 0 {
				cache[key] = 1
				return cache[key]
			} else {
				cache[key] = 0
				return cache[key]
			}
		}

		horizontalCut := 0
		for row := top + 1; row < len(pizza); row++ {
			if apples[top][left]-apples[row][left] == 0 {
				continue
			}

			horizontalCut += numWays(row, left, k-1)
		}

		verticleCut := 0
		for col := left + 1; col < len(pizza[0]); col++ {
			if apples[top][left]-apples[top][col] == 0 {
				continue
			}

			verticleCut += numWays(top, col, k-1)
		}

		cache[key] = (horizontalCut + verticleCut) % mod
		return cache[key]
	}

	return numWays(0, 0, k)
}
