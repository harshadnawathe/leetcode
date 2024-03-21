package reverselinkedlist

// time: O(n)
// space: O(1)
func reverseList(head *ListNode) *ListNode {
	var node, prev *ListNode = head, nil
	for node != nil {
		next := node.Next
		node.Next = prev

		prev = node
		node = next
	}
	return prev
}
