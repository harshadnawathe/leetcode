package findifarraycanbesorted

import (
	"math/bits"
)

func canSortArray(nums []int) bool {
	numBits := map[int]int{}
	for _, num := range nums {
		numBits[num] = bits.OnesCount16(uint16(num))
	}

	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j-1] <= nums[j] {
				continue
			}

			if numBits[nums[j]] != numBits[nums[j-1]] {
				return false
			}

			nums[j], nums[j-1] = nums[j-1], nums[j]

		}
	}

	return true
}
