package numberofoperationstomakenetworkconnected

func makeConnected(n int, connections [][]int) int {
	if len(connections) < n-1 {
		return -1
	}

	graph := makeAdjacencyList(connections)
	visited := make(map[int]struct{})

	var dfs func(int)
	dfs = func(x int) {
		if _, found := visited[x]; found {
			return
		}

		visited[x] = struct{}{}

		for _, next := range graph[x] {
			dfs(next)
		}
	}

	numIslands := 0
	for i := 0; i < n; i++ {
		if _, found := visited[i]; !found {
			numIslands++
			dfs(i)
		}
	}

	return numIslands - 1
}

func makeAdjacencyList(edges [][]int) map[int][]int {
	adjacencyList := make(map[int][]int)
	for _, edge := range edges {
		adjacencyList[edge[0]] = append(adjacencyList[edge[0]], edge[1])
		adjacencyList[edge[1]] = append(adjacencyList[edge[1]], edge[0])
	}
	return adjacencyList
}
