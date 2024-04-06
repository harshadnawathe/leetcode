package containsduplicateii

// time: O(n)
// space: O(n)
func containsNearbyDuplicate(nums []int, k int) bool {
	numIndex := make(map[int]int)

	for i, num := range nums {
		if index, ok := numIndex[num]; ok && i-index <= k {
			return true
		}
		numIndex[num] = i
	}

	return false
}
