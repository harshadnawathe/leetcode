package dividearrayintoarrayswithmaxdifference

import "sort"

func divideArray(nums []int, k int) [][]int {
	parts := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return [][]int{}
		}

		parts = append(parts, nums[i:i+3])
	}

	return parts
}
