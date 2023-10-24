package findlargestvalueineachtreerow

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	q := []*TreeNode{root}
	for len(q) > 0 {
		var next []*TreeNode
		max := math.MinInt32
		for i := 0; i < len(q); i++ {
			node := q[i]
			max = maxOf(max, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		q = next
		result = append(result, max)
	}
	return result
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
