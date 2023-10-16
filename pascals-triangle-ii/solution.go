package pascalstriangleii

func getRow(rowIndex int) []int {
	lastRow := []int{1}
	for r := 1; r <= rowIndex; r++ {
		curr := make([]int, r+1)
		curr[0] = 1
		for c := 1; c < r; c++ {
			curr[c] = lastRow[c-1] + lastRow[c]
		}
		curr[r] = 1

		lastRow = curr
	}
	return lastRow
}
