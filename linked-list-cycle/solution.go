package linkedlistcycle

func hasCycle(head *ListNode) bool {
	return collisionPoint(head) != nil
}

func collisionPoint(node *ListNode) *ListNode {
	if node == nil {
		return nil
	}

	slow, fast := node, node.Next

	for slow != fast {
		slow = slow.Next

		if fast == nil {
			return nil
		}

		fast = fast.Next

		if fast == nil {
			return nil
		}

		fast = fast.Next
	}

	return slow
}
