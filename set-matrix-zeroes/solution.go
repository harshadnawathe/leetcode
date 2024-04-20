package setmatrixzeroes

func setZeroes(matrix [][]int) {
	isZeroRow := make(map[int]bool)
	isZeroColumn := make(map[int]bool)

	for r, row := range matrix {
		for c, val := range row {
			if val == 0 {
				isZeroRow[r] = true
				isZeroColumn[c] = true
			}
		}
	}

	for r, row := range matrix {
		for c := range row {
			if isZeroRow[r] || isZeroColumn[c] {
				matrix[r][c] = 0
			}
		}
	}
}
