package smallestnumberininfiniteset

import "container/heap"

type SmallestInfiniteSet struct {
	smallest int
	mh       *minHeap
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		smallest: 1,
		mh:       newMinHeap(),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if this.mh.Len() != 0 {
		return heap.Pop(this.mh).(int)
	}
	n := this.smallest
	this.smallest += 1
	return n
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.smallest {
		return
	}
	heap.Push(this.mh, num)
}

type minHeap struct {
	h   []int
	set map[int]struct{}
}

func newMinHeap() *minHeap {
	mh := &minHeap{
		h:   []int{},
		set: make(map[int]struct{}),
	}
	heap.Init(mh)
	return mh
}

func (minheap *minHeap) Len() int {
	return len(minheap.h)
}

func (minheap *minHeap) Less(i int, j int) bool {
	return minheap.h[i] < minheap.h[j]
}

func (minheap *minHeap) Swap(i int, j int) {
	minheap.h[i], minheap.h[j] = minheap.h[j], minheap.h[i]
}

func (minheap *minHeap) Push(x any) {
	n := x.(int)
	if _, found := minheap.set[n]; found {
		return
	}
	minheap.set[n] = struct{}{}
	minheap.h = append(minheap.h, n)
}

func (minheap *minHeap) Pop() any {
	old := minheap.h
	n := len(old) - 1
	x := old[n]
	minheap.h = old[0:n]
	delete(minheap.set, x)
	return x
}
