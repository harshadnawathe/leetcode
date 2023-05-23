package kthlargestelementinastream

import "container/heap"

type minHeap []int

func (h *minHeap) Len() int           { return len(*h) }
func (h *minHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeap) Push(x any)         { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n]
	return x
}

type KthLargest struct {
	*minHeap
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{}
	kl.minHeap = &minHeap{}
	kl.k = k

	for _, num := range nums {
		insert(&kl, num)
	}

	return kl
}

func (this *KthLargest) Add(val int) int {
	insert(this, val)
	return min(this)
}

func insert(kl *KthLargest, x int) {
	if kl.Len() == kl.k && x <= min(kl) {
		return
	}

	heap.Push(kl, x)

	if kl.Len() > kl.k {
		heap.Pop(kl)
	}
}

func min(kl *KthLargest) int {
	slice := *(kl.minHeap)
	return slice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
