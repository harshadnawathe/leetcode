package intersectionoftwoarrays

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}

	set1 := makeSet(nums1)
	set2 := makeSet(nums2)

	for num := range set1 {
		if _, ok := set2[num]; ok {
			result = append(result, num)
		}
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
