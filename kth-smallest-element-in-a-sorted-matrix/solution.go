package kthsmallestelementinasortedmatrix

import (
	"container/heap"
)

// Time: O(n^2)
// Space: O(1)
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	low, high := matrix[0][0], matrix[n-1][n-1]
	for low <= high {
		mid := (low + high) / 2
		pos := 0
		result := low
		for i, j := 0, n-1; i < n; i++ {
			for j >= 0 && matrix[i][j] > mid {
				j--
			}

			if j >= 0 {
				pos += 1 + j // number we are looking is on right w
				result = maxOf(result, matrix[i][j])
			} else {
				break
			}
		}
		if pos == k {
			return result
		}
		if pos < k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}

// Time: O(n^2.log(n))
// Space: O(n^2)
func kthSmallest1(matrix [][]int, k int) int {
	h := &minHeap{}
	heap.Init(h)

	for _, row := range matrix {
		for _, val := range row {
			heap.Push(h, val)
		}
	}

	for ; k > 1; k-- {
		heap.Pop(h)
	}
	return heap.Pop(h).(int)
}

type minHeap []int

func (h *minHeap) Len() int               { return len(*h) }
func (h *minHeap) Less(i int, j int) bool { return (*h)[i] < (*h)[j] }
func (h *minHeap) Swap(i int, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeap) Push(x any)             { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() any {
	old := *h
	n := len(old) - 1
	x := old[n]
	*h = old[0:n]
	return x
}
