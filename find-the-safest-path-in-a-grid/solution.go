package findthesafestpathinagrid

import (
	"sort"
)

func maximumSafenessFactor(grid [][]int) int {
	calcMinDistFromTheif(grid)

	maxSafety := maxVal(grid)

	return sort.Search(maxSafety+1, func(safety int) bool {
		return !isPathPresent(grid, safety)
	}) - 1
}

func maxVal(grid [][]int) int {
	result := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			result = max(result, grid[r][c])
		}
	}

	return result
}

func isPathPresent(grid [][]int, safety int) bool {
	destR, destC := len(grid)-1, len(grid[0])-1

	if grid[0][0] < safety || grid[destR][destC] < safety {
		return false
	}

	q := [][2]int{{0, 0}}
	visited := map[[2]int]struct{}{
		{0, 0}: {},
	}

	for len(q) > 0 {
		r, c := q[0][0], q[0][1]
		q = q[1:]

		if r == destR && c == destC {
			return true
		}

		neighbours := [][2]int{{r, c + 1}, {r + 1, c}, {r, c - 1}, {r - 1, c}}

		for _, neighbour := range neighbours {
			nr, nc := neighbour[0], neighbour[1]

			isValid := isValidCell(grid, nr, nc)
			_, isVisited := visited[neighbour]

			if isValid && !isVisited && grid[nr][nc] >= safety {
				visited[neighbour] = struct{}{}
				q = append(q, neighbour)
			}
		}
	}

	return false
}

func calcMinDistFromTheif(grid [][]int) {
	q := [][]int{}

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 1 {
				q = append(q, []int{r, c})
			}
			grid[r][c] -= 1
		}
	}

	for len(q) > 0 {
		nq := len(q)

		for i := 0; i < nq; i++ {
			r, c := q[i][0], q[i][1]

			neighbours := [][]int{{r, c + 1}, {r + 1, c}, {r, c - 1}, {r - 1, c}}

			for _, neighbour := range neighbours {
				nr, nc := neighbour[0], neighbour[1]
				val := grid[r][c]
				if isValidCell(grid, nr, nc) && grid[nr][nc] == -1 {
					q = append(q, neighbour)
					grid[nr][nc] = val + 1
				}
			}
		}

		q = q[nq:]
	}
}

func isValidCell(grid [][]int, r, c int) bool {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return false
	}
	return true
}
