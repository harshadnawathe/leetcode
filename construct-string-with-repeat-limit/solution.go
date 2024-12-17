package constructstringwithrepeatlimit

import "strings"

func repeatLimitedString(s string, repeatLimit int) string {
	charCount := make([]int, 'z'-'a'+1)
	for i := range len(s) {
		charCount[s[i]-'a'] += 1
	}

	var b strings.Builder

	for char := len(charCount) - 1; char >= 0; char-- {
		for charCount[char] > 0 {
			numCharsUsed := min(charCount[char], repeatLimit)
			for range numCharsUsed {
				b.WriteByte(byte('a' + char))
			}
			charCount[char] -= numCharsUsed

			if charCount[char] > 0 {
				nextChar := char - 1
				for nextChar >= 0 && charCount[nextChar] < 1 {
					nextChar--
				}

				if nextChar < 0 {
					break
				}

				b.WriteByte(byte('a' + nextChar))
				charCount[nextChar]--
			}
		}
	}

	return b.String()
}
