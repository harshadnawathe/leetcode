package dividearrayinsetsofkconsecutivenumbers

import "sort"

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	sort.Ints(nums)

	numCount := map[int]int{}
	for _, num := range nums {
		numCount[num]++
	}

	numGroups := 0

	for _, num := range nums {
		if numCount[num] == 0 {
			continue
		}

		numGroups++
		numCount[num]--

		for i := 1; i < k; i++ {
			if numCount[num+i] == 0 {
				return false
			}

			numCount[num+i]--
		}
	}

	return numGroups == len(nums)/k
}
