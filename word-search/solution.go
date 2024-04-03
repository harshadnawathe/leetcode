package wordsearch

func exist(board [][]byte, word string) bool {
	var dfs func(r, c, pos int) bool

	dfs = func(r, c, pos int) bool {
		if pos == len(word) {
			return true
		}

		if r < 0 || r >= len(board) || c < 0 || c >= len(board[0]) {
			return false
		}

		char := board[r][c]
		if char == '$' {
			return false
		}

		board[r][c] = '$'
		defer func() {
			board[r][c] = char
		}()

		if word[pos] != char {
			return false
		}

		return dfs(r+1, c, pos+1) || dfs(r-1, c, pos+1) || dfs(r, c+1, pos+1) || dfs(r, c-1, pos+1)
	}

	for r := range board {
		for c := range board[r] {
			if dfs(r, c, 0) {
				return true
			}
		}
	}

	return false
}
