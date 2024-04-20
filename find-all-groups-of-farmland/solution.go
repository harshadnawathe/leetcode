package findallgroupsoffarmland

func findFarmland(land [][]int) [][]int {
	var farmLandBbx = func(r, c int) []int {
		bbx := []int{r, c, r, c}

		q := [][]int{{r, c}}
		for len(q) > 0 {
			cell := q[0]
			q = q[1:]

			if cell[0] == len(land) || cell[1] == len(land[0]) {
				continue
			}

			if land[cell[0]][cell[1]] == 0 {
				continue
			}

			land[cell[0]][cell[1]] = 0

			bbx[2] = max(bbx[2], cell[0])
			bbx[3] = max(bbx[3], cell[1])

			q = append(q, []int{cell[0] + 1, cell[1]}, []int{cell[0], cell[1] + 1})
		}

		return bbx
	}

	farmLands := [][]int{}
	for r, row := range land {
		for c, cell := range row {
			if cell == 0 {
				continue
			}
			farmLands = append(farmLands, farmLandBbx(r, c))
		}
	}

	return farmLands
}
