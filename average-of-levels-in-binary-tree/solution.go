package averageoflevelsinbinarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
	average := []float64{}

	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		var sum float64
		var nextLevel []*TreeNode

		for _, node := range nodes {
			sum += float64(node.Val)

			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}

			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}

		average = append(average, sum/float64(len(nodes)))
		nodes = nextLevel
	}

	return average
}
