package doubleanumberrepresentedasalinkedlist

// time: O(n)
// space: O(1)
func doubleIt(head *ListNode) *ListNode {
	if head.Val > 4 {
		head = &ListNode{0, head}
	}

	curr := head
	for ; curr.Next != nil; curr = curr.Next {
		if curr.Next.Val > 4 {
			curr.Val = 1 + (curr.Val*2)%10
		} else {
			curr.Val = (curr.Val * 2) % 10
		}
	}

	curr.Val = (curr.Val * 2) % 10

	return head
}

// time: O(n)
// space: O(1)
func doubleIt1(head *ListNode) *ListNode {
	head = reverse(head)

	carry := false
	for curr := head; curr != nil; curr = curr.Next {
		curr.Val *= 2
		if carry {
			curr.Val += 1
		}

		curr.Val, carry = adjustCarry(curr.Val)
	}

	if carry {
		return &ListNode{
			Val:  1,
			Next: reverse(head),
		}
	}

	return reverse(head)
}

func adjustCarry(x int) (int, bool) {
	if x > 9 {
		return x - 10, true
	} else {
		return x, false
	}
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode

	for curr := head; curr != nil; {
		curr.Next, prev, curr = prev, curr, curr.Next
	}

	return prev
}
