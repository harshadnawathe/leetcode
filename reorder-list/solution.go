package reorderlist

// time: O(n)
// space: O(1)
func reorderList(head *ListNode) {
	tail := reverse(midPoint(head))

	for head.Next != nil && tail.Next != nil {
		headRef, tailRef := head, tail

		head = head.Next
		tail = tail.Next
		headRef.Next = tailRef
		tailRef.Next = head
	}
}

func midPoint(head *ListNode) *ListNode {
	slow := head
	for fast := head; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	prev := (*ListNode)(nil)
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}
