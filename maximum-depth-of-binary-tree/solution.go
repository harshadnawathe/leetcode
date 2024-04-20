package maximumdepthofbinarytree

func maxDepth(root *TreeNode) int {
	var maxD func(*TreeNode, int) int
	maxD = func(node *TreeNode, d int) int {
		if node == nil {
			return d
		}
		return max(
			maxD(node.Left, d+1),
			maxD(node.Right, d+1),
		)
	}

	return maxD(root, 0)
}
