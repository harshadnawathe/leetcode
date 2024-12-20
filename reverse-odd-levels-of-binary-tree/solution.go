package reverseoddlevelsofbinarytree

func reverseOddLevels(root *TreeNode) *TreeNode {
	var f func(left, right *TreeNode, depth int)
	f = func(left, right *TreeNode, depth int) {
		if left == nil || right == nil {
			return
		}

		if depth%2 == 1 {
			left.Val, right.Val = right.Val, left.Val
		}

		f(left.Left, right.Right, depth+1)
		f(left.Right, right.Left, depth+1)
	}

	f(root.Left, root.Right, 1)

	return root
}
