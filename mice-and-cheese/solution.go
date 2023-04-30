package miceandcheese

import (
	"container/heap"
	"sort"
)

// Using MaxHeap
// Time: O(nlogn)
// Space: O(n)
func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	maxPoints := 0
	var diff maxHeap
	heap.Init(&diff)
	for i := range reward2 {
		heap.Push(&diff, reward1[i]-reward2[i])
		maxPoints += reward2[i]
	}

	for i := 0; i < k; i++ {
		maxPoints += heap.Pop(&diff).(int)
	}

	return maxPoints
}

type maxHeap []int

func (h *maxHeap) Len() int {
	return len(*h)
}

func (h *maxHeap) Less(i int, j int) bool {
	return (*h)[i] >= (*h)[j]
}

func (h *maxHeap) Swap(i int, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[0:n]
	return x
}

// Using sort
// Time: O(nlog(n))
// Space: O(n)
func miceAndCheese1(reward1 []int, reward2 []int, k int) int {
	maxPoints := 0
	diff := []int{}
	for i := 0; i < len(reward1); i++ {
		diff = append(diff, reward1[i]-reward2[i])
		maxPoints += reward2[i]
	}

	sort.Slice(diff, func(i, j int) bool {
		return diff[i] >= diff[j]
	})

	for i := 0; i < k; i++ {
		maxPoints += diff[i]
	}

	return maxPoints
}
