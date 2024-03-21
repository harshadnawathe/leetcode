package reverselinkedlist

// iterative solution
// time: O(n)
// space: O(1)
func reverseList(head *ListNode) *ListNode {
	var node, prev *ListNode = head, nil
	for node != nil {
		next := node.Next
		node.Next = prev

		prev = node
		node = next
	}
	return prev
}

// // recursive solution
// // // time: O(n)
// // // space: O(n) -- call stack
// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
//
// 	if head.Next == nil {
// 		return head
// 	}
//
// 	tail := head.Next
// 	newHead := reverseList(tail)
// 	tail.Next = head
// 	head.Next = nil
//
// 	return newHead
// }
