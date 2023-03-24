package reorderroutestomakeallpathsleadtothecityzero

func minReorder(n int, connections [][]int) int {
	graph := makeAdjacencyList(connections)

	visited := make(map[int]struct{})

	reorderCount := 0
	var dfs func(int)
	dfs = func(parent int) {
		visited[parent] = struct{}{}

		for _, child := range graph[parent] {
			//child is parent
			if _, found := visited[child.id]; found {
				continue
			}
			if child.dir == down {
				reorderCount++
			}
			dfs(child.id)
		}
	}

	dfs(0)
	return reorderCount
}

type direction int

const (
	unknown direction = iota
	up
	down
)

type city struct {
	id  int
	dir direction
}

func makeAdjacencyList(edges [][]int) map[int][]city {
	graph := make(map[int][]city)
	for _, edge := range edges {
		src, dest := edge[0], edge[1]
		graph[src] = append(graph[src], city{dest, down})
		graph[dest] = append(graph[dest], city{src, up})
	}
	return graph
}
