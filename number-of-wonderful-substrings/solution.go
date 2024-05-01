package numberofwonderfulsubstrings

func wonderfulSubstrings(word string) int64 {
	var count int64
	mask := 0
	maskCount := map[int]int{0: 1}

	for i := 0; i < len(word); i++ {
		mask ^= maskFor(word[i])

		count += int64(maskCount[mask])

		maskCount[mask] += 1

		for b := byte('a'); b <= byte('j'); b++ {
			count += int64(maskCount[mask^maskFor(b)])
		}
	}
	return count
}

func maskFor(b byte) int {
	return 1 << int(b-'a')
}
