package islantperimeter

func islandPerimeter(grid [][]int) int {
	landCount, neighborCount := 0, 0

	for r, row := range grid {
		for c, cell := range row {
			if cell != 0 {
				landCount++

				if r > 0 && grid[r-1][c] != 0 {
					neighborCount++
				}

				if c > 0 && grid[r][c-1] != 0 {
					neighborCount++
				}
			}
		}
	}
	return landCount*4 - neighborCount*2
}
