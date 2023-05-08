package firstcompletelypaintedroworcolumn

func firstCompleteIndex(arr []int, mat [][]int) int {
	coordinate := make(map[int][2]int)
	rows := make([]int, len(mat))
	cols := make([]int, len(mat))

	for r, row := range mat {
		for c, val := range row {
			coordinate[val] = [2]int{r, c}
		}
	}

	for i, val := range arr {
		rc := coordinate[val]

		rows[rc[0]]++
		cols[rc[1]]++

		if rows[rc[0]] == len(mat[0]) || cols[rc[1]] == len(mat) {
			return i
		}
	}

	return len(arr)
}

type set map[int]struct{}

func firstCompleteIndex1(arr []int, mat [][]int) int {
	coordinate := make(map[int][2]int)

	rows := make([]set, len(mat))
	for i := range rows {
		rows[i] = make(set)
	}

	cols := make([]set, len(mat[0]))
	for i := range cols {
		cols[i] = make(set)
	}

	for r, row := range mat {
		for c, val := range row {
			coordinate[val] = [2]int{r, c}

			rows[r][val] = struct{}{}
			cols[c][val] = struct{}{}
		}
	}

	for i, val := range arr {
		rc := coordinate[val]

		delete(rows[rc[0]], val)
		if len(rows[rc[0]]) == 0 {
			return i
		}

		delete(cols[rc[1]], val)
		if len(cols[rc[1]]) == 0 {
			return i
		}
	}

	return len(arr)
}
