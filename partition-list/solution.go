package partitionlist

func partition(head *ListNode, x int) *ListNode {
	ltx := &ListNode{}
	ltxEnd := ltx

	gtex := &ListNode{}
	gtexEnd := gtex

	for node := head; node != nil; node = node.Next {
		if node.Val < x {
			ltxEnd.Next = node
			ltxEnd = ltxEnd.Next
		} else {
			gtexEnd.Next = node
			gtexEnd = gtexEnd.Next
		}
	}

	ltxEnd.Next = gtex.Next
	gtexEnd.Next = nil

	return ltx.Next
}
