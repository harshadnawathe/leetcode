package longestcycleinagraph

func longestCycle(edges []int) int {
	result := -1

	visited := make(map[int]struct{})

	var dfs func(int, map[int]int)
	dfs = func(node int, distance map[int]int) {
		if _, found := visited[node]; found {
			return
		}
		visited[node] = struct{}{}

		next := edges[node]
		if next == -1 {
			return
		}

		if d, found := distance[next]; found {
			//cycle detected
			result = maxOf(result, distance[node]-d+1)
			return
		}

		distance[next] = distance[node] + 1
		dfs(next, distance)
	}

	for i := 0; i < len(edges); i++ {
		dfs(i, map[int]int{i: 1})
	}

	return result
}

func maxOf(lhs, rhs int) int {
	if lhs < rhs {
		return rhs
	}
	return lhs
}
