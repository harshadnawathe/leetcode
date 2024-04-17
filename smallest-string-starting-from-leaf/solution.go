package smalleststringstartingfromleaf

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func smallestFromLeaf(root *TreeNode) string {
	// set smallest to string larger than z so that min can return the first string
	smallest := "~"

	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, curr string) {
		s := string('a'+byte(node.Val)) + curr

		if node.Left == nil && node.Right == nil {
			smallest = min(smallest, s)
			return
		}

		if node.Left != nil {
			dfs(node.Left, s)
		}

		if node.Right != nil {
			dfs(node.Right, s)
		}
	}

	dfs(root, "")

	return smallest
}
