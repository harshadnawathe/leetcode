package matrixdiagonalsum

func diagonalSum(mat [][]int) int {
	sum := 0

	for i := 0; i < len(mat)/2; i++ {
		top, bottom := i, len(mat)-i-1
		left, right := i, len(mat)-i-1
		sum += mat[top][left]
		sum += mat[top][right]
		sum += mat[bottom][left]
		sum += mat[bottom][right]
	}

	if len(mat)&1 == 1 {
		mid := len(mat) / 2
		sum += mat[mid][mid]
	}

	return sum
}

func diagonalSum2(mat [][]int) int {
	sum := 0
	for i := 0; i < len(mat); i++ {
		sum += mat[i][i]
		sum += mat[len(mat)-i-1][i]
	}

	if len(mat)&1 == 1 {
		sum -= mat[len(mat)/2][len(mat)/2]
	}

	return sum
}

func diagonalSum1(mat [][]int) int {
	sum := 0
	for i := 0; i < len(mat); i++ {
		sum += mat[i][i]
	}

	for r, c := 0, len(mat)-1; r < len(mat) && c >= 0; r, c = r+1, c-1 {
		sum += mat[r][c]
	}

	if len(mat)&1 == 1 {
		mid := len(mat) / 2
		sum -= mat[mid][mid]
	}
	return sum
}
