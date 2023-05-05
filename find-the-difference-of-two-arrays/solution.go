package findthedifferenceoftwoarrays

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := makeSet(nums1)
	set2 := makeSet(nums2)

	result := [][]int{{}, {}}
	for n := range set1 {
		if _, found := set2[n]; found {
			continue
		}
		result[0] = append(result[0], n)
	}

	for n := range set2 {
		if _, found := set1[n]; found {
			continue
		}
		result[1] = append(result[1], n)
	}
	return result
}

func makeSet(nums []int) map[int]struct{} {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}
	return set
}
