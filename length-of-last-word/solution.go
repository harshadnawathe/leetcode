package lengthoflastword

// time: O(n), space: O(1)
func lengthOfLastWord(s string) int {
	rbegin := len(s) - 1

	for rbegin >= 0 && s[rbegin] == ' ' {
		rbegin--
	}

	lastWordEnd := rbegin

	for rbegin >= 0 && s[rbegin] != ' ' {
		rbegin--
	}

	return lastWordEnd - rbegin
}
