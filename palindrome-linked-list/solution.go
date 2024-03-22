package palindromelinkedlist

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head.Next
	var prev *ListNode = nil
	oddLen := true

	for fast != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next

		fast = fast.Next
		if fast == nil {
			oddLen = false
			break
		}

		fast = fast.Next
	}

	if oddLen {
		slow = slow.Next
	}

	for prev != nil && slow != nil {
		if prev.Val != slow.Val {
			return false
		}

		prev = prev.Next
		slow = slow.Next
	}

	return true
}
