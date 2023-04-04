package mergeksortedlists

import "container/heap"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	h := &IntHeap{}
	heap.Init(h)
	for _, list := range lists {
		for list != nil {
			heap.Push(h, list.Val)
			list = list.Next
		}
	}

	head := &ListNode{}
	tail := head

	for h.Len() > 0 {
		tail.Next = &ListNode{Val: heap.Pop(h).(int)}
		tail = tail.Next
	}
	return head.Next
}

type IntHeap []int

func (h *IntHeap) Len() int {
	return len(*h)
}

func (h *IntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *IntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
