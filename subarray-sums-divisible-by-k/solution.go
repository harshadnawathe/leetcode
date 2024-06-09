package subarraysumsdivisiblebyk

func subarraysDivByK(nums []int, k int) int {
	modCount := map[int]int{0: 1}

	subarraysCount := 0

	prefixMod := 0

	for _, num := range nums {

		prefixMod = (prefixMod + (num % k) + k) % k

		subarraysCount += modCount[prefixMod]

		modCount[prefixMod] += 1
	}

	return subarraysCount
}
