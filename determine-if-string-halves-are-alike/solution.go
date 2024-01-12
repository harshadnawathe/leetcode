package determineifstringhalvesarealike

// time: O(n)
// space: O(1)
func halvesAreAlike(s string) bool {
	vowelCountA := vowelCount(s, 0, len(s)/2)
	vowelCountB := vowelCount(s, len(s)/2, len(s))

	return vowelCountA == vowelCountB
}

var vowels = map[byte]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'A': {},
	'E': {},
	'I': {},
	'O': {},
	'U': {},
}

func vowelCount(s string, start, end int) int {
	count := 0
	for i := start; i < end; i++ {
		if _, ok := vowels[s[i]]; ok {
			count++
		}
	}
	return count
}
