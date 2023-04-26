package cherrypickup

import "math"

func cherryPickup(grid [][]int) int {
	cache := make(map[[2]pos]float64)
	var numCherry func(pos, pos) float64
	numCherry = func(p1, p2 pos) float64 {
		if !isValid(p1, grid) || !isValid(p2, grid) {
			return math.Inf(-1)
		}

		if isThorn(p1, grid) || isThorn(p2, grid) {
			return math.Inf(-1)
		}

		if isAtBottomRight(p1, grid) {
			return float64(grid[p1.row][p1.col])
		}

		if isAtBottomRight(p2, grid) {
			return float64(grid[p2.row][p2.col])
		}

		key := [2]pos{p1, p2}
		if cached, found := cache[key]; found {
			return cached
		}

		var cherry float64
		if equal(p1, p2) {
			cherry += float64(grid[p1.row][p1.col])
		} else {
			cherry += float64(grid[p1.row][p1.col] + grid[p2.row][p2.col])
		}

		cherry += math.Max(
			math.Max(numCherry(down(p1), down(p2)), numCherry(down(p1), right(p2))),
			math.Max(numCherry(right(p1), down(p2)), numCherry(right(p1), right(p2))),
		)
		cache[key] = cherry
		return cherry
	}

	n := numCherry(pos{0, 0}, pos{0, 0})
	if math.IsInf(n, -1) {
		return 0
	}

	return int(n)
}

type pos struct {
	row, col int
}

func equal(p1, p2 pos) bool { return p1.row == p2.row && p1.col == p2.col }
func right(p pos) pos       { return pos{p.row, p.col + 1} }
func down(p pos) pos        { return pos{p.row + 1, p.col} }
func isValid(p pos, grid [][]int) bool {
	return p.row >= 0 && p.row < len(grid) && p.col >= 0 && p.col < len(grid[0])
}

func isThorn(p pos, grid [][]int) bool {
	return grid[p.row][p.col] == -1
}

func isAtBottomRight(p pos, grid [][]int) bool {
	return p.row == len(grid)-1 && p.col == len(grid[0])-1
}
