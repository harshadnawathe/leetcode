package mergetwosortedlists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	merged := &ListNode{}
	tail := merged

	for list1 != nil && list2 != nil {
		if list2.Val < list1.Val {
			tail.Next = list2
			list2 = list2.Next
		} else {
			tail.Next = list1
			list1 = list1.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
		tail = tail.Next
	}
	if list2 != nil {
		tail.Next = list2
		tail = tail.Next
	}

	return merged.Next
}
