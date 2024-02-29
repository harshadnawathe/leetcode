package evenoddtree

import "errors"

func isEven(x int) bool {
	return x&0x1 == 0
}

func isOdd(x int) bool {
	return x&0x1 != 0
}

type conditionFunc func([]*TreeNode, *TreeNode) bool

func evenLevelCondition(level []*TreeNode, x *TreeNode) bool {
	if !isOdd(x.Val) {
		return false
	}
	if len(level) == 0 {
		return true
	}
	last := level[len(level)-1]
	return last.Val < x.Val
}

func oddLevelCondition(level []*TreeNode, x *TreeNode) bool {
	if !isEven(x.Val) {
		return false
	}
	if len(level) == 0 {
		return true
	}
	last := level[len(level)-1]
	return x.Val < last.Val
}

func conditionFuncFor(level int) conditionFunc {
	if isEven(level) {
		return evenLevelCondition
	}
	return oddLevelCondition
}

func appendIf(cond conditionFunc, a []*TreeNode, x *TreeNode) ([]*TreeNode, error) {
	if !cond(a, x) {
		return nil, errors.New("pre-condition failed")
	}
	return append(a, x), nil
}

// time: O(n) where n is the number of nodes
// space: O(n) where n is the number of nodes
func isEvenOddTree(root *TreeNode) bool {
	if isEven(root.Val) {
		return false
	}

	level := 0
	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		level++
		var nextNodes []*TreeNode

		for _, node := range nodes {
			var err error

			if node.Left != nil {
				nextNodes, err = appendIf(conditionFuncFor(level), nextNodes, node.Left)
				if err != nil {
					return false
				}
			}

			if node.Right != nil {
				nextNodes, err = appendIf(conditionFuncFor(level), nextNodes, node.Right)
				if err != nil {
					return false
				}
			}

		}
		nodes = nextNodes
	}

	return true
}
