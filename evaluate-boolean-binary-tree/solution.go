package evaluatebooleanbinarytree

func evaluateTree(root *TreeNode) bool {
	var evalNonNil func(*TreeNode) bool
	var eval func(*TreeNode, bool) bool

	evalNonNil = func(node *TreeNode) bool {
		switch node.Val {
		case 0:
			return false
		case 1:
			return true
		case 2:
			return eval(node.Left, false) || eval(node.Right, false)
		case 3:
			return eval(node.Left, true) && eval(node.Right, true)
		default:
			panic("bad node value")
		}
	}

	eval = func(node *TreeNode, defaultVal bool) bool {
		if node == nil {
			return defaultVal
		}
		return evalNonNil(node)
	}

	return evalNonNil(root)
}
