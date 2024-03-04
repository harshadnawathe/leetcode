package countofsmallernumbersafterself

func countSmaller(nums []int) []int {
	count := make([]int, len(nums))
	position := make([]int, len(nums))
	for i := range nums {
		position[i] = i
	}

	buffer := make([]int, len(position))
	mid := len(position) / 2
	leftBuff, rightBuff := buffer[:mid], buffer[mid:]

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
			for j < len(aRhs) && nums[aLhs[i]] > nums[aRhs[j]] {
				j++
			}
			count[aLhs[i]] += j
		}

		copy(leftBuff, aLhs)
		copy(rightBuff, aRhs)

		l, r := 0, 0
		for i := range a {
			if l < len(aLhs) && r < len(aRhs) {
				if nums[leftBuff[l]] < nums[rightBuff[r]] {
					a[i] = leftBuff[l]
					l++
				} else {
					a[i] = rightBuff[r]
					r++
				}
			} else if l < len(aLhs) {
				a[i] = leftBuff[l]
				l++
			} else if r < len(aRhs) {
				a[i] = rightBuff[r]
				r++
			}
		}
	}

	mergeSort(position)
	return count
}
