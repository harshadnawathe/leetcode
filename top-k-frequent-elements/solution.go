package topkfrequentelements

import (
	"container/heap"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	sort.Ints(nums)

	var h numCounts
	heap.Init(&h)

	for i := 0; i < len(nums); {
		j := i + 1
		for j < len(nums) && nums[i] == nums[j] {
			j++
		}

		heap.Push(&h, numCount{nums[i], j - i})
		i = j
	}

	result := []int{}
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(&h).(numCount).num)
	}

	return result
}

type numCount struct {
	num, cnt int
}

type numCounts []numCount

func (h *numCounts) Len() int {
	return len(*h)
}

func (h *numCounts) Less(i int, j int) bool {
	return (*h)[j].cnt < (*h)[i].cnt
}

func (h *numCounts) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *numCounts) Push(x any) {
	*h = append(*h, x.(numCount))
}

func (h *numCounts) Pop() any {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n]
	return x
}
