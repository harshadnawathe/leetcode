package scoreafterflippingmatrix

func matrixScore(grid [][]int) int {
	nR, nC := len(grid), len(grid[0])

	score := (1 << (nC - 1)) * nR

	for c := 1; c < nC; c++ {
		cntEqHeader := 0

		for r := 0; r < nR; r++ {
			if grid[r][0] == grid[r][c] {
				cntEqHeader++
			}
		}

		cntEqHeader = max(cntEqHeader, nR-cntEqHeader)

		score += (1 << (nC - c - 1)) * cntEqHeader
	}

	return score
}
