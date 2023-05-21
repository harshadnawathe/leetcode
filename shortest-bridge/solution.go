package shortestbridge

func shortestBridge(grid [][]int) int {
	var nodeOnFirstIsland node
search:
	for r, row := range grid {
		for c, val := range row {
			if val == 1 {
				nodeOnFirstIsland.row = r
				nodeOnFirstIsland.col = c
				break search
			}
		}
	}

	nodes := markAllNodesInIsland(grid, nodeOnFirstIsland, 2)

	distance := 0
	for len(nodes) > 0 {
		next := []node{}
		for _, n := range nodes {
			for _, nextNode := range nextNodes(n) {
				if nextNode.row < 0 || nextNode.row >= len(grid) || nextNode.col < 0 || nextNode.col >= len(grid[0]) {
					continue
				}
				if grid[nextNode.row][nextNode.col] == 1 {
					return distance
				}
				if grid[nextNode.row][nextNode.col] != 0 {
					continue
				}
				grid[nextNode.row][nextNode.col] = -1
				next = append(next, nextNode)
			}
		}
		distance++
		nodes = next
	}
	return distance
}

type node struct {
	row, col int
}

func nextNodes(n node) []node {
	return []node{
		{n.row - 1, n.col},
		{n.row, n.col + 1},
		{n.row + 1, n.col},
		{n.row, n.col - 1},
	}
}

func markAllNodesInIsland(grid [][]int, nodeOnIsland node, marker int) []node {
	nodes := []node{}
	var dfs func(node)
	dfs = func(n node) {
		if n.row < 0 || n.row >= len(grid) || n.col < 0 || n.col >= len(grid[0]) {
			return
		}
		if grid[n.row][n.col] != 1 {
			return
		}
		grid[n.row][n.col] = marker

		nodes = append(nodes, n)

		for _, nextNode := range nextNodes(n) {
			dfs(nextNode)
		}
	}
	dfs(nodeOnIsland)
	return nodes
}
