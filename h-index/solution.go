package hindex

import "sort"

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] > citations[j]
	})

	return sort.Search(len(citations), func(i int) bool {
		return citations[i] <= i
	})
}
