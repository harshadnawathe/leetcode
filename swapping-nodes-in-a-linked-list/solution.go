package swappingnodesinalinkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	handle := &ListNode{Next: head}

	lhs := advance(handle, k)

	slow := handle
	fast := advance(handle, k)
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	rhs := slow

	lhs.Val, rhs.Val = rhs.Val, lhs.Val
	return head
}

func advance(node *ListNode, steps int) *ListNode {
	for steps > 0 && node != nil {
		steps--
		node = node.Next
	}
	return node
}
