package largestlocalvaluesinamatrix

func largestLocal(grid [][]int) [][]int {
	result := [][]int{}

	for i := 1; i < len(grid)-1; i++ {
		rr := []int{}
		for j := 1; j < len(grid[0])-1; j++ {
			rr = append(rr, findLocalMax(grid, i, j))
		}
		result = append(result, rr)
	}

	return result
}

func findLocalMax(grid [][]int, i, j int) int {
	localMax := grid[i][j]

	for r := i - 1; r < i+2; r++ {
		for c := j - 1; c < j+2; c++ {
			localMax = max(localMax, grid[r][c])
		}
	}

	return localMax
}
