package pseudopalindromicpathsinabinarytree

func pseudoPalindromicPaths(root *TreeNode) int {
	count := 0
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, numCount int) {
		numCount ^= 1 << node.Val

		if node.Left != nil {
			dfs(node.Left, numCount)
		}
		if node.Right != nil {
			dfs(node.Right, numCount)
		}

		if node.Left == nil && node.Right == nil {
			if numCount&(numCount-1) == 0 {
				count++
			}
		}
	}
	dfs(root, 0)
	return count
}
