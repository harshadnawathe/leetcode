package largestcolorvalueinadirectedgraph

import (
	"errors"
	"math"
)

func largestPathValue(colors string, edges [][]int) int {
	graph := adjList(edges)
	count := make(map[int][]int)
	for i := 0; i < len(colors); i++ {
		count[i] = make([]int, 26)
	}
	visited := make(map[int]struct{})
	inStack := make(map[int]struct{})

	var dfs func(int) (int, error)
	dfs = func(node int) (int, error) {
		if _, present := inStack[node]; present {
			return 0, errors.New("cycle detected")
		}

		inStack[node] = struct{}{}
		defer delete(inStack, node)

		if _, present := visited[node]; present {
			return count[node][int(colors[node]-'a')], nil
		}
		visited[node] = struct{}{}

		for _, neighbour := range graph[node] {
			if _, err := dfs(neighbour); err != nil {
				return 0, err
			}
			for i := 0; i < 26; i++ {
				count[node][i] = maxOf(count[node][i], count[neighbour][i])
			}
		}

		count[node][int(colors[node]-'a')]++
		return count[node][int(colors[node]-'a')], nil
	}

	max := math.MinInt32
	for i := 0; i < len(colors); i++ {
		x, err := dfs(i)
		if err != nil {
			return -1
		}
		max = maxOf(max, x)
	}

	return max
}

func adjList(edges [][]int) map[int][]int {
	m := make(map[int][]int)
	for _, edge := range edges {
		m[edge[0]] = append(m[edge[0]], edge[1])
	}
	return m
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
