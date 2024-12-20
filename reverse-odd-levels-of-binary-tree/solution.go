package reverseoddlevelsofbinarytree

func reverseOddLevels(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	level := 0

	for len(q) > 0 {
		if level%2 == 1 {
			reverseNodeValues(q)
		}

		for range len(q) {
			node := q[0]
			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		level++
	}

	return root
}

func reverseNodeValues(nodes []*TreeNode) {
	left, right := 0, len(nodes)
	for left < right {
		right--
		nodes[right].Val, nodes[left].Val = nodes[left].Val, nodes[right].Val
		left++
	}
}
