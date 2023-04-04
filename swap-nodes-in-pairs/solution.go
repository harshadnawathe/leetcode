package swapnodesinpairs

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	handle := &ListNode{Next: head}
	curr := handle

	for curr.Next != nil && curr.Next.Next != nil {
		first := curr.Next
		second := curr.Next.Next

		first.Next = second.Next
		second.Next = first
		curr.Next = second

		curr = curr.Next.Next
	}

	return handle.Next
}
