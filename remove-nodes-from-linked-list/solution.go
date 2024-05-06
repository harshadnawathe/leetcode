package removenodesfromlinkedlist

// time: O(n)
// space: O(1)
func removeNodes(head *ListNode) *ListNode {
	head = reverse(head)

	for curr := head; curr != nil; curr = curr.Next {
		for curr.Next != nil && curr.Val > curr.Next.Val {
			curr.Next = curr.Next.Next
		}
	}

	return reverse(head)
}

func reverse(list *ListNode) *ListNode {
	var prev *ListNode

	for curr := list; curr != nil; {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return prev
}

// time: O(n)
// space: O(n)
func removeNodes1(head *ListNode) *ListNode {
	stk := stack{}

	for head != nil {
		for topNode, ok := top(stk); ok && topNode.Val < head.Val; topNode, ok = top(stk) {
			stk = pop(stk)
		}

		stk = push(stk, head)

		head = head.Next
	}

	for i := 0; i < len(stk)-1; i++ {
		stk[i].Next = stk[i+1]
	}

	return stk[0]
}

type stack []*ListNode

func push(stk stack, x *ListNode) stack { return append(stk, x) }
func pop(stk stack) stack               { return stk[:len(stk)-1] }

func top(stk stack) (*ListNode, bool) {
	if len(stk) == 0 {
		return nil, false
	}
	return stk[len(stk)-1], true
}
