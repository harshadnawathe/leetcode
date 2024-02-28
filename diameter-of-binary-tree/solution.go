package diameterofbinarytree

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter int

	var nodeHeight func(*TreeNode) int
	nodeHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftHeight := nodeHeight(node.Left)
		rightHeight := nodeHeight(node.Right)

		diameter = maxOf(diameter, leftHeight+rightHeight)

		return 1 + maxOf(leftHeight, rightHeight)
	}

	_ = nodeHeight(root)

	return diameter
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
