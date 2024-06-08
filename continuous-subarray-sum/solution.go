package continuoussubarraysum

func checkSubarraySum(nums []int, k int) bool {
	modIndex := map[int]int{0: -1}

	prefixMod := 0

	for i, num := range nums {
		prefixMod = (prefixMod + num) % k

		if index, ok := modIndex[prefixMod]; ok {
			if i-index > 1 {
				return true
			}
		} else {
			modIndex[prefixMod] = i
		}
	}

	return false
}
