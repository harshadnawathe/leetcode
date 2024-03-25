package sumofdistances

// time: O(n)
// space: O(n)
func distance(nums []int) []int64 {
	arr := make([]int64, len(nums))

	type sumCount struct {
		sum, count int
	}

	numToSumCount := make(map[int]*sumCount)

	for i, num := range nums {
		if sc, ok := numToSumCount[num]; ok {
			sc.sum += i
			sc.count++
		} else {
			numToSumCount[num] = &sumCount{sum: i, count: 1}
		}

		arr[i] = int64(numToSumCount[num].count*i) - int64(numToSumCount[num].sum)
	}

	numToSumCount = make(map[int]*sumCount)

	for i := len(nums) - 1; i >= 0; i-- {
		num := nums[i]

		if sc, ok := numToSumCount[num]; ok {
			sc.sum += i
			sc.count++
		} else {
			numToSumCount[num] = &sumCount{sum: i, count: 1}
		}

		arr[i] -= int64(numToSumCount[num].count*i) - int64(numToSumCount[num].sum)
	}

	return arr
}
