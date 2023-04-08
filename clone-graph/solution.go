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

	nodes := map[int]*Node{
		node.Val: &Node{Val: node.Val},
	}

	q := queue{node}
	for len(q) > 0 {
		curr := front(q)
		q = remove(q)

		for _, neighbour := range curr.Neighbors {
			if _, ok := nodes[neighbour.Val]; !ok {
				nodes[neighbour.Val] = &Node{Val: neighbour.Val}
				q = add(q, neighbour)
			}
			nodes[curr.Val].Neighbors = append(nodes[curr.Val].Neighbors, nodes[neighbour.Val])
		}
	}

	return nodes[node.Val]
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
