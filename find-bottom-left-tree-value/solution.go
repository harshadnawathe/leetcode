package findbottomlefttreevalue

func findBottomLeftValue(root *TreeNode) int {
	level := []*TreeNode{root}
	var bottomLeft int

	for len(level) > 0 {
		var nextLevel []*TreeNode

		for _, node := range level {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		bottomLeft = level[0].Val
		level = nextLevel
	}

	return bottomLeft
}
