package outofboundarypaths

const mod = 1000000007

// Recursion with memoization
// time: O(row * col * maxMove)
// space: O(row * col * maxMove)
// func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
// 	cache := make(map[[3]int]int)
//
//
// 	var numPaths func(int, int, int) int
// 	numPaths = func(moves int, r,c int) int {
// 		if isOutOfBoundary(m, n, r, c) {
// 			return 1
// 		}
//
// 		if moves == maxMove {
// 			return 0
// 		}
//
// 		key := [3]int{moves, r, c}
// 		if cached, ok := cache[key]; ok {
// 			return cached
// 		}
//
// 		positions := [][]int{
// 			{r - 1, c}, // Up
// 			{r, c + 1}, // Right
// 			{r + 1, c}, // Down
// 			{r, c - 1}, // Left
// 		}
//
// 		pathCount := 0
// 		for _, position := range positions {
// 			pathCount += numPaths(moves+1, position[0], position[1])
// 			pathCount %= mod
// 		}
// 		cache[key] = pathCount
// 		return pathCount
// 	}
//
// 	return numPaths(0, startRow, startColumn)
// }

// Tabular DP
// time: O(row * col * moves)
// space: O(row * col)
func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	numPaths := make([][][]int, 2)
	for i := range numPaths {
		numPaths[i] = make([][]int, m)
		for j := range numPaths[i] {
			numPaths[i][j] = make([]int, n)
		}
	}

	prev, curr := 0, 1
	for moves := maxMove - 1; moves >= 0; moves-- {
		for r := 0; r < m; r++ {
			for c := 0; c < n; c++ {
				positions := [][]int{
					{r - 1, c}, // Up
					{r, c + 1}, // Right
					{r + 1, c}, // Down
					{r, c - 1}, // Left
				}
				numPaths[curr][r][c] = 0
				for _, position := range positions {
					if isOutOfBoundary(m, n, position[0], position[1]) {
						numPaths[curr][r][c] += 1
					} else {
						numPaths[curr][r][c] += numPaths[prev][position[0]][position[1]]
					}
					numPaths[curr][r][c] %= mod
				}
			}
		}
		prev, curr = curr, prev
	}

	return numPaths[prev][startRow][startColumn]
}

func isOutOfBoundary(m, n int, r, c int) bool {
	if r < 0 || r >= m || c < 0 || c >= n {
		return true
	} else {
		return false
	}
}
