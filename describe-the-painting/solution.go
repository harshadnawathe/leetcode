package describethepainting

import (
	"sort"
)

func splitPainting(segments [][]int) [][]int64 {
	colotAt := make(map[int]int64)
	for _, segment := range segments {
		colotAt[segment[0]] += int64(segment[2])
		colotAt[segment[1]] -= int64(segment[2])
	}

	keys := make([]int, 0, len(colotAt))
	for k := range colotAt {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	result := [][]int64{}
	sum := int64(0)

	begin := 0
	for _, end := range keys {
		if sum > 0 {
			result = append(result, []int64{int64(begin), int64(end), sum})
		}
		sum += colotAt[end]
		begin = end
	}

	return result
}
