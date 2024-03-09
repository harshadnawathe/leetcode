package minimumcommonvalue

// time complexity O(n+m)
// space complexity O(1)
func getCommon(nums1 []int, nums2 []int) int {
	first, second := 0, 0

	for first < len(nums1) && second < len(nums2) {
		if nums1[first] < nums2[second] {
			first++
		} else if nums2[second] < nums1[first] {
			second++
		} else {
			return nums1[first]
		}
	}
	return -1
}

// // time complexity O(n + m)
// // space complexity O(n)
// func getCommon(nums1 []int, nums2 []int) int {
// 	set1 := make(map[int]struct{})
// 	for _, num := range nums1 {
// 		set1[num] = struct{}{}
// 	}
//
// 	for _, num := range nums2 {
// 		if _, ok := set1[num]; ok {
// 			return num
// 		}
// 	}
//
// 	return -1
// }
