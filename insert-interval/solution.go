package insertinterval

import (
	"sort"
)

// time: O(n + log(n))
// space: O(1) -- not accounting for result which will be O(n)
func insert(intervals [][]int, newInterval []int) [][]int {
	begin := sort.Search(len(intervals), func(i int) bool {
		return intervals[i][1] >= newInterval[0]
	})

	if begin == len(intervals) {
		return append(intervals, newInterval)
	}

	end := sort.Search(len(intervals), func(i int) bool {
		return intervals[i][0] > newInterval[1]
	})

	if begin != end {
		rangeInterval := []int{intervals[begin][0], intervals[end-1][1]}
		newInterval = merge(newInterval, rangeInterval)
	}

	result := append(intervals[:begin+1], intervals[end:]...)
	result[begin] = newInterval

	return result
}

func merge(interval, other []int) []int {
	interval[0] = minOf(interval[0], other[0])
	interval[1] = maxOf(interval[1], other[1])
	return interval
}

func minOf(lhs, rhs int) int {
	if lhs < rhs {
		return lhs
	}
	return rhs
}

func maxOf(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}
