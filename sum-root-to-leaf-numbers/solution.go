package sumroottoleafnumbers

func sumNumbers(root *TreeNode) int {
	total := 0

	var f func(*TreeNode, int)
	f = func(node *TreeNode, sum int) {
		next := sum*10 + node.Val

		if node.Left == nil && node.Right == nil {
			total += next
			return
		}

		if node.Left != nil {
			f(node.Left, next)
		}

		if node.Right != nil {
			f(node.Right, next)
		}
	}

	f(root, 0)

	return total
}
