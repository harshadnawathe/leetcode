package validpalindrome

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		for left < right && !isAlphanumberic(s[left]) {
			left++
		}

		for right > left && !isAlphanumberic(s[right]) {
			right--
		}

		if toLowerCase(s[left]) != toLowerCase(s[right]) {
			return false
		}
		left++
		right--
	}

	return true
}

func isAlphanumberic(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9')
}

func toLowerCase(char byte) byte {
	if char >= 'A' && char <= 'Z' {
		return 'a' + char - 'A'
	}
	return char
}
