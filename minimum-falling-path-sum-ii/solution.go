package minimumfallingpathsumii

import "math"

func minFallingPathSum(grid [][]int) int {
	var minSum func(int, int) int

	cache := make(map[[2]int]int)

	minSum = func(row, col int) int {
		if row == len(grid)-1 {
			return grid[row][col]
		}

		key := [2]int{row, col}

		if cached, ok := cache[key]; ok {
			return cached
		}

		nextMinSum := math.MaxInt32

		for c := 0; c < col; c++ {
			nextMinSum = min(nextMinSum, minSum(row+1, c))
		}

		for c := col + 1; c < len(grid[0]); c++ {
			nextMinSum = min(nextMinSum, minSum(row+1, c))
		}

		result := nextMinSum + grid[row][col]

		cache[key] = result

		return result
	}

	result := math.MaxInt32

	for c := 0; c < len(grid[0]); c++ {
		result = min(result, minSum(0, c))
	}

	return result
}
