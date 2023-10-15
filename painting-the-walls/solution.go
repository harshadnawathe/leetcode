package paintingthewalls

import "math"

// tabular dp
// O(n^2) where n = len(cost)
func paintWalls(cost []int, time []int) int {
	minCost := make([][]int, len(cost)+1)
	for i := range minCost {
		minCost[i] = make([]int, len(cost)+1)
	}

	for i := 1; i <= len(cost); i++ {
		minCost[len(cost)][i] = math.MaxInt32
	}

	for i := len(cost) - 1; i >= 0; i-- {
		for n := 0; n <= len(cost); n++ {
			minCost[i][n] = minOf(
				cost[i]+minCost[i+1][maxOf(0, n-time[i]-1)],
				minCost[i+1][n],
			)
		}
	}

	return minCost[0][len(cost)]
}

// dp with memoization
// O(n^2) where n = len(cost)
func paintWalls1(cost []int, time []int) int {
	cache := make(map[[2]int]int)

	var paintWallsMinCost func(int, int) int
	paintWallsMinCost = func(i, n int) int {
		if n <= 0 {
			return 0
		}
		if i == len(cost) {
			return math.MaxInt32
		}
		key := [2]int{i, n}
		if cached, ok := cache[key]; ok {
			return cached
		}
		minCost := minOf(
			cost[i]+paintWallsMinCost(i+1, n-time[i]-1),
			paintWallsMinCost(i+1, n),
		)
		cache[key] = minCost
		return minCost
	}

	return paintWallsMinCost(0, len(cost))
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}

func minOf(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}
