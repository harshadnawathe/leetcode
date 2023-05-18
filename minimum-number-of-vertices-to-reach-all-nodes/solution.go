package minimumnumberofverticestoreachallnodes

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	numIncomingEdges := make([]int, n)
	for _, edge := range edges {
		numIncomingEdges[edge[1]]++
	}

	vertices := []int{}
	for i, n := range numIncomingEdges {
		if n == 0 {
			vertices = append(vertices, i)
		}
	}
	return vertices
}
