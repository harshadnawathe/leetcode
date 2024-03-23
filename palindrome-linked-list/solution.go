package palindromelinkedlist

// time: O(n)
// space: O(1)
func isPalindrome(head *ListNode) bool {
	rev := reverse(midPoint(head))

	for head != nil && rev != nil {
		if head.Val != rev.Val {
			return false
		}
		head = head.Next
		rev = rev.Next
	}
	return true
}

func midPoint(head *ListNode) *ListNode {
	slow := head
	for fast := head; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// time: O(n)
// space: O(1)
// func isPalindrome(head *ListNode) bool {
// 	slow, fast := head, head.Next
// 	var prev *ListNode = nil
// 	oddLen := true
//
// 	for fast != nil {
// 		next := slow.Next
// 		slow.Next = prev
// 		prev = slow
// 		slow = next
//
// 		fast = fast.Next
// 		if fast == nil {
// 			oddLen = false
// 			break
// 		}
//
// 		fast = fast.Next
// 	}
//
// 	if oddLen {
// 		slow = slow.Next
// 	}
//
// 	for prev != nil && slow != nil {
// 		if prev.Val != slow.Val {
// 			return false
// 		}
//
// 		prev = prev.Next
// 		slow = slow.Next
// 	}
//
// 	return true
// }
