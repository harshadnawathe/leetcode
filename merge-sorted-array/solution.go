package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {
	head := len(nums1) - 1
	i1, i2 := m-1, n-1

	for i1 >= 0 && i2 >= 0 {
		if nums1[i1] > nums2[i2] {
			nums1[head] = nums1[i1]
			i1--
		} else {
			nums1[head] = nums2[i2]
			i2--
		}
		head--
	}

	for i2 >= 0 {
		nums1[head] = nums2[i2]
		i2--
		head--
	}
}
