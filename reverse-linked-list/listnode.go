package reverselinkedlist

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func listOf(xs ...int) *ListNode {
	handle := &ListNode{}
	head := handle

	for _, x := range xs {
		head.Next = &ListNode{Val: x}
		head = head.Next
	}

	return handle.Next
}

func equal(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

func (l *ListNode) String() string {
	var b strings.Builder
	b.WriteByte('[')
	for l != nil {
		b.WriteString(fmt.Sprintf("%d ", l.Val))
		l = l.Next
	}
	b.WriteByte(']')
	return b.String()
}
