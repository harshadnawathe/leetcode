package sumofdistancesintree

func sumOfDistancesInTree(n int, edges [][]int) []int {
	count := make([]int, n)
	distSum := make([]int, n)
	tree := makeAdjacencyList(edges)

	var distSumNode0 func(int, int)

	distSumNode0 = func(root, prev int) {
		for _, next := range tree[root] {
			if next == prev {
				continue
			}
			distSumNode0(next, root)
			count[root] += count[next]
			distSum[root] += distSum[next] + count[next]
		}
		count[root] += 1
	}

	var distSumRootShift func(int, int)

	distSumRootShift = func(node, prev int) {
		for _, next := range tree[node] {
			if next == prev {
				continue
			}
			distSum[next] = distSum[node] - count[next] + n - count[next]
			distSumRootShift(next, node)
		}
	}

	distSumNode0(0, -1)
	distSumRootShift(0, -1)

	return distSum
}

func makeAdjacencyList(edges [][]int) map[int][]int {
	g := make(map[int][]int)

	for _, edge := range edges {
		g[edge[0]] = append(g[edge[0]], edge[1])
		g[edge[1]] = append(g[edge[1]], edge[0])
	}

	return g
}
