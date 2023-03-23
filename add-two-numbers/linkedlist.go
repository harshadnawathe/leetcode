package addtwonumbers

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func makeLinkedList(nums ...int) *ListNode {
	handle := &ListNode{}
	tail := handle

	for _, num := range nums {
		tail.Next = &ListNode{num, nil}
		tail = tail.Next
	}

	return handle.Next
}

func equal(list1, list2 *ListNode) bool {

	for list1 != nil && list2 != nil {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}

	return list1 == nil && list2 == nil

}

func (list *ListNode) String() string {
	var b strings.Builder

	b.WriteByte('[')
	for list != nil {
		b.WriteString(fmt.Sprintf(" %d", list.Val))
		list = list.Next
	}
	b.WriteByte(']')

	return b.String()
}
