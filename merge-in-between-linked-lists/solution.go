package mergeinbetweenlinkedlists

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	i := 0
	node1 := list1
	for i < a-1 {
		node1 = node1.Next
		i++
	}
	aPrev := node1

	for i < b+1 {
		node1 = node1.Next
		i++
	}
	bNext := node1

	node2 := list2
	for node2.Next != nil {
		node2 = node2.Next
	}

	aPrev.Next = list2
	node2.Next = bNext

	return list1
}
