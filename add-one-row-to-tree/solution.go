package addonerowtotree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time: O(n), space: O(n) where n number of nodes
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}

	currLevel := []*TreeNode{root}

	for depth-1 > 1 && len(currLevel) > 0 {
		var nextLevel []*TreeNode
		for _, node := range currLevel {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		depth--
		currLevel = nextLevel
	}

	for _, node := range currLevel {
		node.Left = &TreeNode{Val: val, Left: node.Left}
		node.Right = &TreeNode{Val: val, Right: node.Right}
	}

	return root
}
