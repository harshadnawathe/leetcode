package findthelongestvalidobstaclecourseateachposition

import "sort"

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	result := make([]int, 0, len(obstacles))
	lis := make([]int, 0, len(obstacles))

	for _, obstacle := range obstacles {
		// upper bound
		x := sort.Search(len(lis), func(i int) bool {
			return lis[i] > obstacle
		})
		if x == len(lis) {
			lis = append(lis, obstacle)
		} else {
			lis[x] = obstacle
		}
		result = append(result, x+1)
	}

	return result
}
