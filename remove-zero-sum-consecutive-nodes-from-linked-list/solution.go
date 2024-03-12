package removezerosumconsecutivenodesfromlinkedlist

// time: O(n), space: O(n)
func removeZeroSumSublists(head *ListNode) *ListNode {
	handle := &ListNode{Next: head}
	curr := handle
	prefixSum := 0
	prefixSumNode := make(map[int]*ListNode)

	for curr != nil {
		prefixSum += curr.Val

		if prevNode, ok := prefixSumNode[prefixSum]; ok {
			curr := prevNode.Next
			p := prefixSum + curr.Val

			for p != prefixSum {
				delete(prefixSumNode, p)
				curr = curr.Next
				p += curr.Val
			}
			prevNode.Next = curr.Next

		} else {
			prefixSumNode[prefixSum] = curr
		}
		curr = curr.Next
	}

	return handle.Next
}
