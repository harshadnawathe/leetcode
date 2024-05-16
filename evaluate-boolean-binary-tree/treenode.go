package evaluatebooleanbinarytree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) String() string {
	if node == nil {
		return "nil"
	}
	return fmt.Sprintf("{%v %v %v}", node.Val, node.Left, node.Right)
}

func treeOf(vals ...any) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}

	i := 0
	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		next := []*TreeNode{}
		for _, node := range nodes {
			if i+1 < len(vals) {
				i++
				if vals[i] != nil {
					left := &TreeNode{Val: vals[i].(int)}
					node.Left = left
					next = append(next, left)
				}
			}

			if i+1 < len(vals) {
				i++
				if vals[i] != nil {
					right := &TreeNode{Val: vals[i].(int)}
					node.Right = right
					next = append(next, right)
				}
			}
		}

		nodes = next
	}

	return root
}
