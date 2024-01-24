package pseudopalindromicpathsinabinarytree

func pseudoPalindromicPaths(root *TreeNode) int {
	count := 0
	var dfs func(*TreeNode, map[int]int)
	dfs = func(node *TreeNode, numCount map[int]int) {
		numCount[node.Val]++
		defer func() { numCount[node.Val]-- }()

		if node.Left != nil {
			dfs(node.Left, numCount)
		}
		if node.Right != nil {
			dfs(node.Right, numCount)
		}

		if node.Left == nil && node.Right == nil {
			if isPseudoPalindromic(numCount) {
				count++
			}
		}
	}
	dfs(root, map[int]int{})
	return count
}

func isPseudoPalindromic(numCount map[int]int) bool {
	numOfOddCounts := 0
	for i := 1; i < 10; i++ {
		if numCount[i]%2 == 1 {
			numOfOddCounts++
		}
		if numOfOddCounts > 1 {
			return false
		}
	}
	return true
}
