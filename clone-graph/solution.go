package clonegraph

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	q := queue{node}
	nodes := make(map[int]*Node)
	visited := make(map[int]struct{})

	for len(q) > 0 {
		curr := front(q)
		q = remove(q)

		if _, ok := visited[curr.Val]; ok {
			continue
		}
		visited[curr.Val] = struct{}{}
		if _, ok := nodes[curr.Val]; !ok {
			nodes[curr.Val] = &Node{Val: curr.Val}
		}

		for _, neighbour := range curr.Neighbors {
			if _, ok := nodes[neighbour.Val]; !ok {
				nodes[neighbour.Val] = &Node{Val: neighbour.Val}
			}
			nodes[curr.Val].Neighbors = append(nodes[curr.Val].Neighbors, nodes[neighbour.Val])
			q = add(q, neighbour)
		}

	}

	return nodes[1]
}

type queue []*Node

func add(q queue, node *Node) queue {
	return append(q, node)
}

func remove(q queue) queue {
	return q[1:]
}

func front(q queue) *Node {
	return q[0]
}
