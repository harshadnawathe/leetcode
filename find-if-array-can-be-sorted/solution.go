package findifarraycanbesorted

import (
	"cmp"
)

func canSortArray(nums []int) bool {
	numBits := map[int]int{}
	for _, num := range nums {
		numBits[num] = bitsCount(num)
	}

	isBitCountSame := func(i, j int) bool {
		return numBits[nums[i]] == numBits[nums[j]]
	}

	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j-1] > nums[j] && isBitCountSame(j, j-1); j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

	return isSorted(nums)
}

func bitsCount(n int) int {
	count := 0

	for i, mask := 0, 1; i <= 8; i++ {

		if mask&n != 0 {
			count++
		}

		mask <<= 1
	}

	return count
}

func isSorted[T cmp.Ordered](a []T) bool {
	for i := range len(a) - 1 {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}
