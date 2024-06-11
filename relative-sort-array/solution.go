package relativesortarray

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	numCount := make(map[int]int)

	for _, num := range arr1 {
		numCount[num] += 1
	}

	result := make([]int, 0, len(arr1))

	for _, num := range arr2 {
		for numCount[num] > 0 {
			result = append(result, num)
			numCount[num] -= 1
		}
	}

	n := len(result)

	for num, count := range numCount {
		for count > 0 {
			result = append(result, num)
			count--
		}
	}

	sort.Ints(result[n:])

	return result
}
