package maximumsubsequencescore

import (
	"container/heap"
	"sort"
)

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	var pairs []struct{ F, S int64 }
	for i := 0; i < len(nums1); i++ {
		pairs = append(pairs, struct {
			F, S int64
		}{int64(nums1[i]), int64(nums2[i])})
	}

	// sort pairs in decreasing order to maximize the initial score
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[j].S < pairs[i].S
	})

	var mh minHeap
	heap.Init(&mh)

	// calculate initial score based on first k pairs
	var topKSum int64
	for i := 0; i < k; i++ {
		topKSum += pairs[i].F
		heap.Push(&mh, pairs[i].F)
	}

	max := topKSum * int64(pairs[k-1].S)

	// adjust score for remaining w pairs
	for i := k; i < len(nums1); i++ {
		// add new element and remove smalled element from sum
		topKSum += int64(pairs[i].F) - heap.Pop(&mh).(int64)
		heap.Push(&mh, pairs[i].F)

		// calculate max
		max = maxOf(max, topKSum*pairs[i].S)
	}

	return max
}

type minHeap []int64

func (h *minHeap) Len() int           { return len(*h) }
func (h *minHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int64)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[:n]
	return x
}

func maxOf(lhs, rhs int64) int64 {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
