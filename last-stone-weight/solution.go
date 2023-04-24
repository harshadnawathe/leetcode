package laststoneweight

import (
	"container/heap"
)

func lastStoneWeight(stones []int) int {
	maxHeap := stonePile(stones)
	heap.Init(&maxHeap)
	for len(maxHeap) >= 2 {
		first := heap.Pop(&maxHeap).(int)
		second := heap.Pop(&maxHeap).(int)

		if first == second {
			continue
		}

		heap.Push(&maxHeap, first-second)
	}

	if len(maxHeap) == 0 {
		return 0
	}
	return maxHeap[0]
}

type stonePile []int

// Len is the number of elements in the collection.
func (sp *stonePile) Len() int {
	return len(*sp)
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
func (sp *stonePile) Less(i int, j int) bool {
	a := *sp
	return a[i] >= a[j]
}

// Swap swaps the elements with indexes i and j.
func (sp *stonePile) Swap(i int, j int) {
	a := *sp
	a[i], a[j] = a[j], a[i]
}

func (sp *stonePile) Push(x any) {
	*sp = append(*sp, x.(int))
}

func (sp *stonePile) Pop() any {
	old := *sp
	n := len(old) - 1
	x := old[n]
	*sp = old[0:n]
	return x
}
