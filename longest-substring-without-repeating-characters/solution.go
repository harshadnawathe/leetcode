package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {
	result := 0
	index := make(map[byte]int)

	for i, j := 0, 0; i < len(s); i++ {
		if x, found := index[s[i]]; found {
			j = x + 1
		}
		index[s[i]] = i
		result = maxOf(result, i-j+1)
	}
	return result
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
