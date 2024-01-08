package rangesumofbst

// time: O(n) where n is number of nodes in the tree
// space: O(n)
func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	stk := stack{root}
	for len(stk) > 0 {
		node := top(stk)
		stk = pop(stk)
		comparison := compareInRange(low, high, node.Val)
		if comparison == 1 {
			if node.Left != nil {
				stk = push(stk, node.Left)
			}
		} else if comparison == -1 {
			if node.Right != nil {
				stk = push(stk, node.Right)
			}
		} else {
			sum += node.Val
			if node.Left != nil {
				stk = push(stk, node.Left)
			}
			if node.Right != nil {
				stk = push(stk, node.Right)
			}
		}
	}
	return sum
}

func compareInRange(low, high int, x int) int {
	if high < low {
		panic("bad range")
	}
	if x < low {
		return -1
	} else if x > high {
		return 1
	} else {
		return 0
	}
}

type stack []*TreeNode

func push(stk stack, x *TreeNode) stack {
	return append(stk, x)
}

func top(stk stack) *TreeNode {
	return stk[len(stk)-1]
}

func pop(stk stack) stack {
	return stk[:len(stk)-1]
}
