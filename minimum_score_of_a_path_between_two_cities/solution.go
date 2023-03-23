package minimumscoreofapathbetweentwocities

import "math"

func minScore(n int, roads [][]int) int {
	result := math.MaxInt32

	graph := makeAdjacencyList(roads)

	visited := make(map[int]struct{})

	var dfs func(int)
	dfs = func(n int) {
		if _, found := visited[n]; found {
			return
		}

		visited[n] = struct{}{}

		for _, next := range graph[n] {
			result = minOf(result, next.score)
			dfs(next.n)
		}
	}

	dfs(1)

	return result
}

func minOf(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

type city struct {
	n     int
	score int
}

func makeAdjacencyList(edges [][]int) map[int][]city {
	adjacencyList := make(map[int][]city)

	for _, edge := range edges {
		adjacencyList[edge[0]] = append(adjacencyList[edge[0]], city{edge[1], edge[2]})
		adjacencyList[edge[1]] = append(adjacencyList[edge[1]], city{edge[0], edge[2]})
	}

	return adjacencyList
}
