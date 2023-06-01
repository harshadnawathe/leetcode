package shortestpathinbinarymatrix

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)

	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}

	grid[0][0] = 1
	q := [][]int{{0, 0, 1}}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]

		if p[0] == n-1 && p[1] == n-1 {
			return p[2]
		}

		for _, dir := range dirs {
			next := dir(p[0], p[1], p[2])

			r, c := next[0], next[1]

			if r < 0 || r >= n || c < 0 || c >= n {
				continue
			}

			if grid[r][c] == 1 {
				continue
			}

			grid[r][c] = 1

			q = append(q, next)
		}
	}

	return -1
}

type dirFunc func(int, int, int) []int

var dirs = []dirFunc{
	func(r, c, d int) []int { return []int{r - 1, c, d + 1} },
	func(r, c, d int) []int { return []int{r - 1, c + 1, d + 1} },
	func(r, c, d int) []int { return []int{r, c + 1, d + 1} },
	func(r, c, d int) []int { return []int{r + 1, c + 1, d + 1} },
	func(r, c, d int) []int { return []int{r + 1, c, d + 1} },
	func(r, c, d int) []int { return []int{r + 1, c - 1, d + 1} },
	func(r, c, d int) []int { return []int{r, c - 1, d + 1} },
	func(r, c, d int) []int { return []int{r - 1, c - 1, d + 1} },
}
