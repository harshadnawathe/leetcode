package kthsmallestprimefraction

import "container/heap"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	fractions := minHeapFraction{}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			heap.Push(&fractions, fraction{
				arr[i],
				arr[j],
				float32(arr[i]) / float32(arr[j]),
			})
		}
	}

	for ; k > 1; k-- {
		heap.Pop(&fractions)
	}

	f := fractions[0]

	return []int{f.N, f.D}
}

type fraction struct {
	N, D int
	V    float32
}

type minHeapFraction []fraction

func (h minHeapFraction) Len() int           { return len(h) }
func (h minHeapFraction) Less(i, j int) bool { return h[i].V < h[j].V }
func (h minHeapFraction) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeapFraction) Push(x any)        { *h = append(*h, x.(fraction)) }
func (h *minHeapFraction) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]

	*h = old[:n-1]

	return x
}
