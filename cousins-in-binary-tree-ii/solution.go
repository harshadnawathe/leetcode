package cousinsinbinarytreeii

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func replaceValueInTree(root *TreeNode) *TreeNode {
	root.Val = 0
	level := []*TreeNode{root}

	for len(level) > 0 {
		nextLevel := []*TreeNode{}
		sum := 0
		for _, node := range level {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
				sum += node.Left.Val
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
				sum += node.Right.Val
			}
		}

		if len(nextLevel) != 0 {
			for _, node := range level {
				childrenSum := 0
				if node.Left != nil {
					childrenSum += node.Left.Val
				}
				if node.Right != nil {
					childrenSum += node.Right.Val
				}

				cousinSum := sum - childrenSum
				if node.Left != nil {
					node.Left.Val = cousinSum
				}
				if node.Right != nil {
					node.Right.Val = cousinSum
				}
			}
		}

		level = nextLevel
	}

	return root
}
