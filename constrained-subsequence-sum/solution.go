package constrainedsubsequencesum

import (
	"container/heap"
)

func constrainedSubsetSum(nums []int, k int) int {
	h := &maxHeap{{nums[0], 0}}
	heap.Init(h)

	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		for (i - (*h)[0].Index) > k {
			heap.Pop(h)
		}

		top := (*h)[0]
		currSum := nums[i] + maxOf(0, top.Val)
		sum = maxOf(sum, currSum)
		heap.Push(h, num{currSum, i})
	}

	return sum
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}

type num struct {
	Val, Index int
}

type maxHeap []num

// Len is the number of elements in the collection.
func (h *maxHeap) Len() int {
	return len(*h)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//   - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//   - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (h *maxHeap) Less(i int, j int) bool {
	return (*h)[j].Val < (*h)[i].Val
}

// Swap swaps the elements with indexes i and j.
func (h *maxHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(num))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
