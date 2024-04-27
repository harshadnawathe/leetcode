package minimumfallingpathsum

func minFallingPathSum(matrix [][]int) int {
	rN, cN := len(matrix), len(matrix[0])

	for r := rN - 2; r >= 0; r-- {
		matrix[r][0] = matrix[r][0] + min(
			matrix[r+1][0],
			matrix[r+1][1],
		)

		for c := 1; c < cN-1; c++ {
			matrix[r][c] = matrix[r][c] + min(
				matrix[r+1][c-1],
				matrix[r+1][c],
				matrix[r+1][c+1],
			)
		}

		matrix[r][cN-1] = matrix[r][cN-1] + min(
			matrix[r+1][cN-2],
			matrix[r+1][cN-1],
		)
	}

	minSum := matrix[0][0]
	for _, sum := range matrix[0][1:] {
		minSum = min(minSum, sum)
	}

	return minSum
}
