package removenthnodefromendoflist

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	for i := 0; i < n; i++ {
		if fast == nil {
			panic("index out of bound")
		}
		fast = fast.Next
	}

	p := &ListNode{Next: head}
	for fast != nil {
		fast = fast.Next
		p = p.Next
	}

	if p.Next == head {
		return head.Next
	}

	p.Next = p.Next.Next

	return head
}
