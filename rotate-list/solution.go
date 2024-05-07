package rotatelist

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	fast := head
	size := 1
	for k > 0 {
		k--
		if fast.Next != nil {
			fast = fast.Next
			size++
		} else {
			k %= size
			fast = head
			size = 1
		}
	}

	slow := head

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	fast.Next = head
	head = slow.Next
	slow.Next = nil

	return head
}
