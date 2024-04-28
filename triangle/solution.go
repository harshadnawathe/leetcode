package triangle

func minimumTotal(triangle [][]int) int {

	for r := len(triangle) - 2; r >= 0; r-- {
		for c := 0; c < len(triangle[r]); c++ {
			triangle[r][c] = triangle[r][c] + min(triangle[r+1][c], triangle[r+1][c+1])
		}
	}

	return triangle[0][0]
}
