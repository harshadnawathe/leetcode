package determineiftwostringsareclose

import "sort"

// complexity: time O(N), space O(1)
func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	cc1 := newCharCount(word1)
	cc2 := newCharCount(word2)

	for i := range cc1 {
		if cc1[i] != cc2[i] && (cc1[i] == 0 || cc2[i] == 0) {
			return false
		}
	}

	sort.Ints(cc1)
	sort.Ints(cc2)

	for i := range cc1 {
		if cc1[i] != cc2[i] {
			return false
		}
	}

	return true
}

type charCount []int

func newCharCount(word string) charCount {
	cc := make(charCount, 26)
	for i := 0; i < len(word); i++ {
		cc[word[i]-'a']++
	}
	return cc
}
