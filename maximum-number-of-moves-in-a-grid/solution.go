package maximumnumberofmovesinagrid

func maxMoves(grid [][]int) int {
	cache := make(map[[2]int]int)

	var f func(int, int) int
	f = func(r, c int) int {
		key := [2]int{r, c}

		if cached, ok := cache[key]; ok {
			return cached
		}

		maxMoves := 0

		for _, next := range nextPos(r, c) {
			if canMove(grid, r, c, next[0], next[1]) {
				maxMoves = max(maxMoves, 1+f(next[0], next[1]))
			}
		}

		cache[key] = maxMoves

		return maxMoves
	}

	maxMoves := 0
	for r := range len(grid) {
		maxMoves = max(maxMoves, f(r, 0))
	}

	return maxMoves
}

func length(grid [][]int) int {
	return len(grid)
}

func width(grid [][]int) int {
	return len(grid[0])
}

func isOnGrid(grid [][]int, r, c int) bool {
	if r >= 0 && r < length(grid) && c >= 0 && c < width(grid) {
		return true
	}
	return false
}

func canMove(grid [][]int, sr, sc int, tr, tc int) bool {
	if !isOnGrid(grid, tr, tc) {
		return false
	}

	if grid[tr][tc] <= grid[sr][sc] {
		return false
	}

	return true
}

func nextPos(r, c int) [][2]int {
	return [][2]int{
		{r - 1, c + 1},
		{r, c + 1},
		{r + 1, c + 1},
	}
}
