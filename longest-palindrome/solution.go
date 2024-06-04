package longestpalindrome

func longestPalindrome(s string) int {
	byteCount := make(map[byte]int)
	for i := range len(s) {
		byteCount[s[i]] += 1
	}

	palindromLen := 0
	oddLen := false

	for _, count := range byteCount {
		if count&0x1 == 0 {
			palindromLen += count
		} else {
			palindromLen += (count - 1)
			oddLen = true
		}
	}

	if oddLen {
		return 1 + palindromLen
	}

	return palindromLen
}
