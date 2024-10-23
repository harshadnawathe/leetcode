package kthlargestsuminabinarytree

import "container/heap"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	q := []*TreeNode{root}

	mh := minHeap{int64(root.Val)}

	heap.Init(&mh)

	for len(q) > 0 {

		var levelSum int64

		for n, N := 0, len(q); n < N; n++ {
			node := top(q)
			q = pop(q)

			if node.Left != nil {
				q = push(q, node.Left)
				levelSum += int64(node.Left.Val)
			}

			if node.Right != nil {
				q = push(q, node.Right)
				levelSum += int64(node.Right.Val)
			}

		}

		if len(q) > 0 {
			heap.Push(&mh, levelSum)
		}

		for mh.Len() > k {
			heap.Pop(&mh)
		}

	}

	if mh.Len() < k {
		return -1
	}

	return mh[0]
}

func push[T any](q []T, x T) []T {
	return append(q, x)
}

func pop[T any](q []T) []T {
	return q[1:]
}

func top[T any](q []T) T {
	return q[0]
}

type minHeap []int64

// Len is the number of elements in the collection.
func (mh *minHeap) Len() int {
	return len(*mh)
}

func (mh *minHeap) Less(i int, j int) bool {
	return (*mh)[i] < (*mh)[j]
}

// Swap swaps the elements with indexes i and j.
func (mh *minHeap) Swap(i int, j int) {
	(*mh)[i], (*mh)[j] = (*mh)[j], (*mh)[i]
}

func (mh *minHeap) Push(x any) {
	*mh = append(*mh, x.(int64))
}

func (mh *minHeap) Pop() any {
	old := *mh
	n := len(old)
	x := old[n-1]

	*mh = old[:n-1]
	return x
}
