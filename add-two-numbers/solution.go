package addtwonumbers

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head

	for l1 != nil && l2 != nil {
		tail.Next = &ListNode{
			Val: l1.Val + l2.Val,
		}

		tail = tail.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		tail.Next = &ListNode{Val: l1.Val}
		tail = tail.Next
		l1 = l1.Next
	}

	for l2 != nil {
		tail.Next = &ListNode{Val: l2.Val}
		tail = tail.Next
		l2 = l2.Next
	}

	for curr := head.Next; curr != nil; curr = curr.Next {
		if curr.Val < 10 {
			continue
		}

		curr.Val -= 10
		if curr.Next != nil {
			curr.Next.Val += 1
		} else {
			curr.Next = &ListNode{Val: 1}
		}
	}

	return head.Next
}
