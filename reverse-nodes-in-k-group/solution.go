package reversenodesinkgroup

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	handle := &ListNode{Next: head}
	curr := handle
	for {
		last := advance(curr, k)
		if last == nil {
			break
		}
		first := curr.Next
		curr.Next = reverse(first, last.Next)
		curr = first
	}
	return handle.Next
}

func advance(node *ListNode, n int) *ListNode {
	for i := 0; i < n && node != nil; i++ {
		node = node.Next
	}
	return node
}

func reverse(start, end *ListNode) *ListNode {
	var curr, prev *ListNode = start, end
	for curr != end {
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	return prev
}
