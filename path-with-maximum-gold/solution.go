package pathwithmaximumgold

func getMaximumGold(grid [][]int) int {
	var dfs func(int, int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			return 0
		}

		if grid[r][c] == 0 {
			return 0
		}

		gold := grid[r][c]

		grid[r][c] = 0
		defer func() {
			grid[r][c] = gold
		}()

		result := 0
		result = max(result, dfs(r-1, c))
		result = max(result, dfs(r, c+1))
		result = max(result, dfs(r+1, c))
		result = max(result, dfs(r, c-1))

		return gold + result
	}

	maxGold := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			maxGold = max(maxGold, dfs(r, c))
		}
	}

	return maxGold
}
