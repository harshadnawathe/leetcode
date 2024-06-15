package ipo

import (
	"container/heap"
	"sort"
)

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	pc := make([][]int, len(profits))
	for i := range pc {
		pc[i] = []int{profits[i], capital[i]}
	}

	sort.Slice(pc, func(i, j int) bool {
		return pc[i][1] < pc[j][1]
	})

	h := maxHeap{}

	i := 0

	for _ = range k {
		for i < len(pc) && pc[i][1] <= w {
			heap.Push(&h, pc[i][0])
			i++
		}

		if len(h) > 0 {
			x := heap.Pop(&h).(int)
			w += x
		}
	}

	return w
}

type maxHeap []int

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Less(i int, j int) bool {
	return m[i] > m[j]
}

func (m maxHeap) Swap(i int, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *maxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *maxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]

	*m = old[:n-1]

	return x
}
