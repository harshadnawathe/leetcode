package numberofclosedislands

const (
	land  = 0
	water = 1
)

var neighbourFunc = map[string]func(int, int) (int, int){
	"north": func(r, c int) (int, int) { return r - 1, c },
	"east":  func(r, c int) (int, int) { return r, c + 1 },
	"south": func(r, c int) (int, int) { return r + 1, c },
	"west":  func(r, c int) (int, int) { return r, c - 1 },
}

func closedIsland(grid [][]int) int {

	var dfs func(int, int) bool
	dfs = func(r, c int) bool {
		if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
			return false
		}

		if grid[r][c] == water {
			return true
		}
		grid[r][c] = water

		isClosed := true
		for _, next := range neighbourFunc {
			isClosed = dfs(next(r, c)) && isClosed
		}
		return isClosed
	}

	var numClosed int
	for r, row := range grid {
		for c := range row {
			if grid[r][c] != water && dfs(r, c) {
				numClosed++
			}
		}
	}

	return numClosed
}
