package maximumtwinsumofalinkedlist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	mid2 := midPoint(head)
	mid1 := reverse(head, mid2)

	max := 0
	for mid1 != nil && mid2 != nil {
		max = maxOf(max, mid1.Val+mid2.Val)
		mid1 = mid1.Next
		mid2 = mid2.Next
	}
	return max
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}

func midPoint(head *ListNode) *ListNode {
	slow, fast := head, head

	for {
		if fast == nil {
			break
		}
		fast = fast.Next
		if fast == nil {
			break
		}
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}

func reverse(first, last *ListNode) *ListNode {
	curr := first
	var prev *ListNode
	for curr != last {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
