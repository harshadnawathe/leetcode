package kthlargestelementinanarray

import (
	"container/heap"
)

func findKthLargest(nums []int, k int) int {
	h := intHeap(nums[:k])

	heap.Init(&h)

	for _, num := range nums[k:] {
		if num > h[0] {
			h[0] = num
			heap.Fix(&h, 0)
		}
	}

	return h[0]
}

type intHeap []int

func (h intHeap) Len() int               { return len(h) }
func (h intHeap) Less(i int, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i int, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x any)            { *h = append(*h, x.(int)) }

func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
