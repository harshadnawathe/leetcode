package surroundedregions

func solve(board [][]byte) {
	var visit func(int, int)

	visited := make(map[[2]int]struct{})

	isVisited := func(r, c int) bool {
		_, found := visited[[2]int{r, c}]
		return found
	}

	visit = func(r, c int) {
		if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
			return
		}

		if board[r][c] == 'X' {
			return
		}

		if isVisited(r, c) {
			return
		}

		visited[[2]int{r, c}] = struct{}{}

		visit(r+1, c)
		visit(r, c+1)
		visit(r, c-1)
		visit(r-1, c)
	}

	for c := 0; c < len(board[0]); c++ {
		visit(0, c)
		visit(len(board)-1, c)
	}

	for r := 1; r < len(board)-1; r++ {
		visit(r, 0)
		visit(r, len(board[0])-1)
	}

	for r, row := range board {
		for c, cell := range row {
			if cell == 'O' && !isVisited(r, c) {
				board[r][c] = 'X'
			}
		}
	}
}
