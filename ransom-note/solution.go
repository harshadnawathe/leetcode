package ransomnote

// time: O(m + n), space: O(1)
func canConstruct(ransomNote string, magazine string) bool {
	charCount := make([]int, 26)

	for _, char := range magazine {
		charCount[char-'a']++
	}

	for _, char := range ransomNote {
		if charCount[char-'a'] == 0 {
			return false
		}
		charCount[char-'a']--
	}

	return true
}
