package leafsimilartrees

// time: O(N) | space: O(N)
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafValues1 := leafValues(root1)
	leafValues2 := leafValues(root2)

	if len(leafValues1) != len(leafValues2) {
		return false
	}

	for i := 0; i < len(leafValues1); i++ {
		if leafValues1[i] != leafValues2[i] {
			return false
		}
	}

	return true
}

// time: O(N) | space: O(N)
func leafValues(root *TreeNode) []int {
	values := []int{}
	stk := []*TreeNode{root}

	for len(stk) > 0 {
		node := stk[len(stk)-1] // top
		stk = stk[:len(stk)-1]  // pop
		if node.Left != nil && node.Right != nil {
			stk = append(stk, node.Left, node.Right)
		} else if node.Left != nil {
			stk = append(stk, node.Left)
		} else if node.Right != nil {
			stk = append(stk, node.Right)
		} else {
			// both children are nil - it is a leaf node
			values = append(values, node.Val)
		}
	}

	return values
}
