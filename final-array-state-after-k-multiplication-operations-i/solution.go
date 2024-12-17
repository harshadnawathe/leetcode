package finalarraystateafterkmultiplicationoperationsi

import (
	"container/heap"
)

func getFinalState(nums []int, k int, multiplier int) []int {
	mh := newMinHeapByRef(nums)
	heap.Init(mh)

	for range k {
		nums[mh.refs[0]] *= multiplier
		heap.Fix(mh, 0)
	}

	return nums
}

type minHeapByRef struct {
	nums []int
	refs []int
}

func (mh *minHeapByRef) Len() int {
	return len(mh.nums)
}

func (mh *minHeapByRef) Less(i int, j int) bool {
	if mh.nums[mh.refs[i]] == mh.nums[mh.refs[j]] {
		return mh.refs[i] < mh.refs[j]
	}

	return mh.nums[mh.refs[i]] < mh.nums[mh.refs[j]]
}

func (mh *minHeapByRef) Swap(i int, j int) {
	mh.refs[i], mh.refs[j] = mh.refs[j], mh.refs[i]
}

func (mh *minHeapByRef) Push(x any) {
	panic("not implemented") // TODO: Implement
}

func (mh *minHeapByRef) Pop() any {
	panic("not implemented") // TODO: Implement
}

func newMinHeapByRef(nums []int) *minHeapByRef {
	refs := make([]int, len(nums))

	for i := range len(nums) {
		refs[i] = i
	}

	return &minHeapByRef{
		nums: nums,
		refs: refs,
	}
}
