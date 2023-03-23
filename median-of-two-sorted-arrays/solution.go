package medianoftwosortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	iN, jN := len(nums1), len(nums2)
	kN := medianPos(iN, jN) + 1

	nums := make([]int, 0, 2)
	for i, j, k := 0, 0, 0; k < kN; k++ {
		var n int
		if i < iN && j < jN {
			if nums1[i] < nums2[j] {
				n = nums1[i]
				i++
			} else {
				n = nums2[j]
				j++
			}
		} else if i < iN {
			n = nums1[i]
			i++
		} else if j < jN {
			n = nums2[j]
			j++
		}

		if k == kN-1 || k == kN-2 {
			nums = append(nums, n)
		}
	}

	if len(nums) == 0 {
		return -1
	}

	if (iN+jN)%2 == 0 {
		return float64(nums[0]+nums[1]) / 2.0
	} else {
		return float64(nums[len(nums)-1])
	}

}

func medianPos(iN, jN int) int {
	switch n := iN + jN; n {
	case 0:
		return -1
	case 1:
		return 0
	default:
		return n / 2
	}
}
