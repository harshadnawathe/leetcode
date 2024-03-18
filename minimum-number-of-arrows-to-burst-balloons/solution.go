package minimumnumberofarrowstoburstballoons

import "sort"

// time: O(n log n)
// space: O(1)
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	arrows := 1
	end := points[0][1]
	for i := 1; i < len(points); i++ {
		if points[i][0] > end {
			arrows++
			end = points[i][1]
		} else {
			end = minOf(end, points[i][1])
		}
	}
	return arrows
}

func minOf(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}
