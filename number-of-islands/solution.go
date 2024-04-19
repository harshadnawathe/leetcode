package numberofislands

func numIslands(grid [][]byte) int {
	var visit func(int, int)
	visit = func(r, c int) {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			return
		}

		if grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'

		visit(r-1, c)
		visit(r, c-1)
		visit(r, c+1)
		visit(r+1, c)
	}

	numIslands := 0
	for r, row := range grid {
		for c, cell := range row {
			if cell == '0' {
				continue
			}
			visit(r, c)
			numIslands++
		}
	}

	return numIslands
}
