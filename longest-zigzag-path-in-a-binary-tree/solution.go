package longestzigzagpathinabinarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *TreeNode) int {
	result := 0

	var zigZag func(*TreeNode, bool, int)
	zigZag = func(node *TreeNode, goLeft bool, count int) {
		if node == nil {
			return
		}

		if result < count {
			result = count
		}

		if goLeft {
			zigZag(node.Left, false, count+1)
			zigZag(node.Right, true, 1)
		} else {
			zigZag(node.Left, false, 1)
			zigZag(node.Right, true, count+1)
		}
	}

	zigZag(root, false, 0)
	zigZag(root, true, 0)

	return result
}
