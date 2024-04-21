package findifpathexistsingraph

// time: O(n) where n nodes in graph
// space: O(n)
func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	visited := make(map[int]struct{})
	q := []int{source}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = struct{}{}

		if node == destination {
			return true
		}

		q = append(q, graph[node]...)
	}

	return false
}
