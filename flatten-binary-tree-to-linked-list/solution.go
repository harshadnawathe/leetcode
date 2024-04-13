package flattenbinarytreetolinkedlist

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(tree *TreeNode) {

	for tree != nil {
		if tree.Left != nil {

			x := tree.Left
			for x.Right != nil {
				x = x.Right
			}

			x.Right = tree.Right
			tree.Right = tree.Left
			tree.Left = nil
		}
		tree = tree.Right
	}
}
