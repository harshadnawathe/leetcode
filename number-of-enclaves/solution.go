package numberofenclaves

const (
	water = 0
	land  = 1
)

var neighbourFunc = map[string]func(int, int) (int, int){
	"north": func(r, c int) (int, int) { return r - 1, c },
	"east":  func(r, c int) (int, int) { return r, c + 1 },
	"south": func(r, c int) (int, int) { return r + 1, c },
	"west":  func(r, c int) (int, int) { return r, c - 1 },
}

func numEnclaves(grid [][]int) int {

	var dfs func(int, int)
	dfs = func(r, c int) {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			return
		}
		if grid[r][c] == water {
			return
		}
		grid[r][c] = water

		for _, neighbour := range neighbourFunc {
			dfs(neighbour(r, c))
		}
	}

	for c := 0; c < len(grid[0]); c++ {
		dfs(0, c)
		dfs(len(grid)-1, c)
	}

	for r := 0; r < len(grid); r++ {
		dfs(r, 0)
		dfs(r, len(grid[0])-1)
	}

	result := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell == land {
				result++
			}
		}
	}
	return result
}
