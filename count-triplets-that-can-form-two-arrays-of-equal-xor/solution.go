package counttripletsthatcanformtwoarraysofequalxor

// time: O(n^3)
// space: O(1)
func countTriplets(arr []int) int {
	count := 0

	for start := 0; start < len(arr)-1; start++ {
		xOrA := 0

		for mid := start + 1; mid < len(arr); mid++ {
			xOrA ^= arr[mid-1]

			xOrB := 0
			for end := mid; end < len(arr); end++ {
				xOrB ^= arr[end]

				if xOrA == xOrB {
					count++
				}
			}
		}
	}

	return count
}
