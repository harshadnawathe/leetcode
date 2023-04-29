package reversenodesinevenlengthgroups

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	curr := &ListNode{Next: head}
	for i := 1; curr != nil; i++ {
		groupEnd, length := next(curr, i)

		if length&1 == 1 {
			curr = groupEnd
			continue
		}

		revFirst := curr.Next
		var revLast *ListNode = nil
		if groupEnd != nil {
			revLast = groupEnd.Next
		}

		curr.Next = reverse(revFirst, revLast)

		curr = revFirst
	}
	return head
}

func next(node *ListNode, n int) (*ListNode, int) {
	if node == nil {
		return nil, 0
	}
	i := 0
	for ; i < n; i++ {
		node = node.Next
		if node == nil {
			break
		}
	}
	return node, i
}

func reverse(first, last *ListNode) *ListNode {
	var prev, curr *ListNode = nil, first
	for curr != last {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	if first != nil {
		first.Next = last
	}
	return prev
}
