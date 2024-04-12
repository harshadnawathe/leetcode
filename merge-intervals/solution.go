package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	m := 0
	for n := 1; n < len(intervals); n++ {
		if overlap(intervals[m], intervals[n]) {
			intervals[m][1] = max(intervals[m][1], intervals[n][1])
		} else {
			m++
			intervals[m] = intervals[n]
		}
	}

	return intervals[:m+1]
}

func overlap(a, b []int) bool {
	return a[0] <= b[1] && a[1] >= b[0]
}
