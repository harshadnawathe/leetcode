package sudokusolver

const (
	symbolEmpty = byte('.')
	symbolFirst = byte('1')
	symbolLast  = byte('9')
)

const (
	puzzleHeight = 9
	puzzleWidth  = 9
	gridHeight   = 3
	gridWidth    = 3
)

func solveSudoku(board [][]byte) {
	var solve func([][2]int) bool
	solve = func(emptyCells [][2]int) bool {
		if len(emptyCells) == 0 {
			return true
		}

		next := emptyCells[0]

		for opt := range options(next[0], next[1], board) {
			board[next[0]][next[1]] = opt
			if solve(emptyCells[1:]) {
				return true
			}
			board[next[0]][next[1]] = symbolEmpty
		}
		return false
	}

	solve(emptyCells(board))
}

func options(row, col int, board [][]byte) map[byte]struct{} {
	opts := make(map[byte]struct{})
	for b := symbolFirst; b <= symbolLast; b++ {
		opts[b] = struct{}{}
	}

	for c := 0; c < puzzleWidth; c++ {
		if value := board[row][c]; value != symbolEmpty {
			delete(opts, board[row][c])
		}
	}

	for r := 0; r < puzzleHeight; r++ {
		if value := board[r][col]; value != symbolEmpty {
			delete(opts, board[r][col])
		}
	}

	gridRowOffset := gridHeight * (row / gridHeight)
	gridColOffset := gridWidth * (col / gridWidth)
	for r := 0; r < gridHeight; r++ {
		for c := 0; c < gridWidth; c++ {
			if value := board[gridRowOffset+r][gridColOffset+c]; value != symbolEmpty {
				delete(opts, value)
			}
		}
	}
	return opts
}

func emptyCells(board [][]byte) [][2]int {
	result := [][2]int{}
	for r := 0; r < puzzleHeight; r++ {
		for c := 0; c < puzzleWidth; c++ {
			if board[r][c] == symbolEmpty {
				result = append(result, [2]int{r, c})
			}
		}
	}
	return result
}
