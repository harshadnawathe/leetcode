package minimumcosttohirekworkers

import (
	"container/heap"
	"math"
	"sort"
)

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	wq := make([]wageQuality, len(quality))

	for i := range quality {
		wq[i] = wageQuality{
			R: float64(wage[i]) / float64(quality[i]),
			Q: quality[i],
		}
	}

	sort.Slice(wq, func(i, j int) bool {
		return wq[i].R < wq[j].R
	})

	total := math.Inf(1)

	totalQ := 0
	heapQ := maxHeap{}
	for i := range wq {
		totalQ += wq[i].Q

		heap.Push(&heapQ, wq[i].Q)

		if len(heapQ) > k {
			totalQ -= heap.Pop(&heapQ).(int)
		}

		if len(heapQ) == k {
			total = min(total, float64(totalQ)*wq[i].R)
		}
	}

	return total
}

type wageQuality struct {
	R float64
	Q int
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]

	*h = old[:n-1]

	return x
}
