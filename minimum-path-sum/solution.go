package minimumpathsum

func minPathSum(grid [][]int) int {
	minSum := make([][]int, len(grid))
	for i := range minSum {
		minSum[i] = make([]int, len(grid[0]))
	}

	minSum[0][0] = grid[0][0]

	for c := 1; c < len(grid[0]); c++ {
		minSum[0][c] = minSum[0][c-1] + grid[0][c]
	}

	for r := 1; r < len(grid); r++ {
		minSum[r][0] = minSum[r-1][0] + grid[r][0]
	}

	for r := 1; r < len(grid); r++ {
		for c := 1; c < len(grid[0]); c++ {
			minSum[r][c] = grid[r][c] + minOf(minSum[r-1][c], minSum[r][c-1])
		}
	}

	return minSum[len(grid)-1][len(grid[0])-1]
}

func minOf(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}
