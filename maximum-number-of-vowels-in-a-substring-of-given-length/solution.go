package maximumnumberofvowelsinasubstringofgivenlength

func maxVowels(s string, k int) int {
	count := 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			count++
		}
	}

	max := count
	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			count++
		}
		if isVowel(s[i-k]) {
			count--
		}
		if max < count {
			max = count
		}
	}
	return max
}

var vowlels = map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}}

func isVowel(b byte) bool {
	_, found := vowlels[b]
	return found
}
