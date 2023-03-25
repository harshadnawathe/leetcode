package countunreachablepairsofnodesinanundirectedgraph

func countPairs(n int, edges [][]int) int64 {
	graph := makeAdjacencyList(edges)

	visited := make(map[int]struct{})

	var dfs func(int) int
	dfs = func(node int) int {
		if _, found := visited[node]; found {
			return 0
		}

		visited[node] = struct{}{}
		count := 1
		for _, next := range graph[node] {
			count += dfs(next)
		}
		return count
	}

	var pairsUnreachable int64
	nodesRemaining := n
	for i := 0; i < n; i++ {
		clusterSize := dfs(i)
		pairsUnreachable += int64(clusterSize) * int64(nodesRemaining-clusterSize)
		nodesRemaining -= clusterSize
	}

	return pairsUnreachable
}

func makeAdjacencyList(edges [][]int) map[int][]int {
	adj := make(map[int][]int)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		adj[edge[1]] = append(adj[edge[1]], edge[0])
	}
	return adj
}
