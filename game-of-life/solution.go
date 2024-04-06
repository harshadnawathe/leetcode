package gameoflife

func gameOfLife(board [][]int) {
	for r := range board {
		for c := range board[r] {
			board[r][c] = nextState(board, r, c)
		}
	}

	for r := range board {
		for c := range board[r] {
			board[r][c] >>= 1
		}
	}
}

const (
	dead       = 0
	live       = 3
	deadToLive = 2
	liveToDead = 1
)

func isCellAlive(board [][]int, r, c int) bool {
	return board[r][c]&0x1 != dead
}

var delta = [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func numberOfLiveNeighbours(board [][]int, r, c int) int {
	count := 0
	for _, d := range delta {
		nr := r + d[0]
		nc := c + d[1]

		if nr < 0 || nr >= len(board) || nc < 0 || nc >= len(board[0]) {
			continue
		}

		if isCellAlive(board, nr, nc) {
			count++
		}
	}
	return count
}

func nextState(board [][]int, r, c int) int {
	numLiveNeighbours := numberOfLiveNeighbours(board, r, c)
	isAlive := isCellAlive(board, r, c)

	if isAlive {
		if numLiveNeighbours < 2 || numLiveNeighbours > 3 {
			return liveToDead
		} else {
			return live
		}
	} else {
		if numLiveNeighbours == 3 {
			return deadToLive
		} else {
			return dead
		}
	}
}
