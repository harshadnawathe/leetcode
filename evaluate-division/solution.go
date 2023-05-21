package evaluatedivision

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := makeGraph(equations, values)

	results := make([]float64, 0, len(queries))
	for _, query := range queries {
		results = append(results, solve(graph, query[0], query[1]))
	}
	return results
}

type adjacentNode struct {
	node  string
	value float64
}

func makeGraph(equations [][]string, values []float64) map[string][]adjacentNode {
	graph := make(map[string][]adjacentNode)

	for i, equation := range equations {
		graph[equation[0]] = append(graph[equation[0]], adjacentNode{equation[1], values[i]})
		graph[equation[1]] = append(graph[equation[1]], adjacentNode{equation[0], 1 / values[i]})
	}

	return graph
}

func solve(graph map[string][]adjacentNode, start, end string) float64 {
	if _, found := graph[start]; !found {
		return -1
	}

	visited := make(map[string]struct{})
	var dfs func(string, float64) float64
	dfs = func(start string, acc float64) float64 {
		if start == end {
			return acc
		}
		visited[start] = struct{}{}

		for _, next := range graph[start] {
			if _, found := visited[next.node]; found {
				continue
			}
			if solution := dfs(next.node, acc*next.value); solution != -1 {
				return solution
			}
		}
		return -1
	}

	return dfs(start, 1.0)
}
