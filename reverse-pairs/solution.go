package reversepairs

// // time: O(n^2)
// // space: O(1)
// func reversePairs(nums []int) int {
// 	var count int
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] > 2*nums[j] {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

// time: O(nlogn)
// space: O(n)
func reversePairs(nums []int) int {
	var count int

	ma := make([]int, len(nums))
	mid := len(ma) / 2
	lhs, rhs := ma[:mid], ma[mid:]

	var mergeSort func([]int)
	mergeSort = func(a []int) {
		if len(a) < 2 {
			return
		}
		mid := len(a) / 2
		aLhs := a[:mid]
		aRhs := a[mid:]

		mergeSort(aLhs)
		mergeSort(aRhs)

		j := 0
		for i := range aLhs {
			for j < len(aRhs) && aLhs[i] > 2*aRhs[j] {
				j++
			}
			count += j
		}

		copy(lhs, aLhs)
		copy(rhs, aRhs)

		li, ri := 0, 0
		for i := range a {
			if li < len(aLhs) && ri < len(aRhs) {
				if lhs[li] < rhs[ri] {
					a[i] = lhs[li]
					li++
				} else {
					a[i] = rhs[ri]
					ri++
				}
			} else if li < len(aLhs) {
				a[i] = lhs[li]
				li++
			} else if ri < len(aRhs) {
				a[i] = rhs[ri]
				ri++
			}
		}
	}

	mergeSort(nums)

	return count
}
