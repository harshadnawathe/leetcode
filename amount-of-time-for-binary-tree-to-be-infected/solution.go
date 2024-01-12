package amountoftimeforbinarytreetobeinfected

// time: O(n)
// space: O(n) where n is number of nodes in the tree
func amountOfTime(root *TreeNode, start int) int {
	graph := adjacencyList(root)
	nodesInfected := map[int]struct{}{start: {}}
	nodesToInfect := graph[start]

	time := 0
	for len(nodesToInfect) > 0 {
		var nextNodesToInfect []int
		for _, node := range nodesToInfect {
			for _, neighbour := range graph[node] {
				if _, infected := nodesInfected[neighbour]; !infected {
					nextNodesToInfect = append(nextNodesToInfect, neighbour)
				}
			}
			nodesInfected[node] = struct{}{}
		}
		nodesToInfect = nextNodesToInfect
		time++
	}

	return time
}

func adjacencyList(root *TreeNode) map[int][]int {
	graph := map[int][]int{}
	stk := []*TreeNode{root}

	for len(stk) > 0 {
		node := stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		if node.Left != nil && node.Right != nil {
			graph[node.Val] = append(graph[node.Val], node.Left.Val, node.Right.Val)
			graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
			graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
			stk = append(stk, node.Left, node.Right)
		} else if node.Left != nil {
			graph[node.Val] = append(graph[node.Val], node.Left.Val)
			graph[node.Left.Val] = append(graph[node.Left.Val], node.Val)
			stk = append(stk, node.Left)
		} else if node.Right != nil {
			graph[node.Val] = append(graph[node.Val], node.Right.Val)
			graph[node.Right.Val] = append(graph[node.Right.Val], node.Val)
			stk = append(stk, node.Right)
		}

	}
	return graph
}
