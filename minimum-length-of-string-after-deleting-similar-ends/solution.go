package minimumlengthofstringafterdeletingsimilarends

func minimumLength(s string) int {
	left, right := 0, len(s)-1
	for left < right && s[left] == s[right] {
		for left+1 < right && s[left] == s[left+1] {
			left++
		}

		for left < right-1 && s[right] == s[right-1] {
			right--
		}

		left++
		right--
	}
	return right - left + 1
}
