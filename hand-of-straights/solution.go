package handofstraights

import "container/heap"

func isNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	set := newNumSet()

	for _, num := range hand {
		add(set, num)
	}

	for set.Len() > 0 {
		num := minimum(set)

		remove(set, num)

		for i := 1; i < groupSize; i++ {
			if !contains(set, num+i) {
				return false
			}
			remove(set, num+i)
		}
	}

	return true
}

type CountPos struct {
	Count, Pos int
}

type numSet struct {
	numCountPos map[int]*CountPos
	minHeap     []int
}

func newNumSet() *numSet {
	return &numSet{
		numCountPos: make(map[int]*CountPos),
		minHeap:     []int{},
	}
}

func add(ns *numSet, num int) {
	if cp, ok := ns.numCountPos[num]; ok {
		cp.Count += 1
	} else {
		heap.Push(ns, num)
	}
}

func remove(ns *numSet, num int) {
	if cp, ok := ns.numCountPos[num]; ok {
		if cp.Count > 1 {
			cp.Count -= 1
		} else {
			heap.Remove(ns, cp.Pos)
		}
	}
}

func contains(ns *numSet, num int) bool {
	_, ok := ns.numCountPos[num]
	return ok
}

func minimum(ns *numSet) int {
	return ns.minHeap[0]
}

func (ns *numSet) Len() int {
	return len(ns.minHeap)
}

func (ns *numSet) Less(i int, j int) bool {
	return ns.minHeap[i] < ns.minHeap[j]
}

func (ns *numSet) Swap(i int, j int) {
	ns.minHeap[i], ns.minHeap[j] = ns.minHeap[j], ns.minHeap[i]
	ns.numCountPos[ns.minHeap[j]].Pos = j
	ns.numCountPos[ns.minHeap[i]].Pos = i
}

func (ns *numSet) Push(x any) {
	num := x.(int)

	ns.numCountPos[num] = &CountPos{1, len(ns.minHeap)}
	ns.minHeap = append(ns.minHeap, num)
}

func (ns *numSet) Pop() any {
	old := ns.minHeap
	n := len(old)

	x := old[n-1]

	ns.minHeap = old[:n-1]
	delete(ns.numCountPos, x)

	return x
}
