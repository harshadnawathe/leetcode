package reverselinkedlistii

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	handle := &ListNode{Next: head}

	node := handle
	for left > 1 {
		node = node.Next
		left--
		right--
	}

	begin := node

	for right > 0 {
		node = node.Next
		right--
	}

	end := node

	begin.Next = reverse(begin.Next, end.Next)

	return handle.Next
}

func reverse(begin, end *ListNode) *ListNode {
	var prev *ListNode

	for curr := begin; curr != end; {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	begin.Next = end

	return prev
}
