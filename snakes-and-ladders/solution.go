package snakeandladders

func snakesAndLadders(board [][]int) int {
	isLeftToRight := true
	pos := 1
	destination := make(map[int]int)

	for r := len(board) - 1; r >= 0; r-- {
		if isLeftToRight {
			for c := 0; c < len(board[0]); c, pos = c+1, pos+1 {
				if board[r][c] != -1 {
					destination[pos] = board[r][c]
				}
			}
		} else {
			for c := len(board[0]) - 1; c >= 0; c, pos = c-1, pos+1 {
				if board[r][c] != -1 {
					destination[pos] = board[r][c]
				}
			}
		}

		isLeftToRight = !isLeftToRight
	}

	q := [][]int{{1, 0}}
	target := len(board) * len(board[0])
	visited := make(map[int]struct{})

	for len(q) > 0 {
		pos, steps := q[0][0], q[0][1]
		q = q[1:]

		if _, ok := visited[pos]; ok {
			continue
		}

		visited[pos] = struct{}{}

		if pos == target {
			return steps
		}

		for i := 1; i <= 6 && pos+i <= target; i++ {
			if dest, ok := destination[pos+i]; ok {
				q = append(q, []int{dest, steps + 1})
			} else {
				q = append(q, []int{pos + i, steps + 1})
			}
		}
	}

	return -1
}
