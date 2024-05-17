package deleteleaveswithagivenvalue

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var canDelete func(*TreeNode) bool

	canDelete = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		delLeft, delRight := canDelete(node.Left), canDelete(node.Right)

		if delLeft && delRight {
			node.Left, node.Right = nil, nil
			return node.Val == target
		} else if delLeft {
			node.Left = nil
			return false
		} else if delRight {
			node.Right = nil
			return false
		} else {
			return false
		}
	}

	if canDelete(root) {
		return nil
	}

	return root
}
