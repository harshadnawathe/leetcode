package middleofthelinkedlist

// time complexity O(n)
// space complexity O(1)
func middleNode(head *ListNode) *ListNode {

	slow, fast := head, head.Next

	for fast != nil {

		slow = slow.Next

		fast = fast.Next

		if fast == nil {
			break
		}

		fast = fast.Next

	}

	return slow
}
