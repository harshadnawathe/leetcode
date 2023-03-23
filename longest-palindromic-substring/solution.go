package longestpalindromicsubstring

func longestPalindrome(s string) string {
	var max string
	for i := 0; i < len(s); i++ {
		evenLen := palindromAround(s, i, i+1)
		oddLen := palindromAround(s, i, i)

		max = maxLenOf3(max, evenLen, oddLen)
	}
	return max
}

func palindromAround(s string, mid1, mid2 int) string {

	for mid1 >= 0 && mid2 < len(s) && s[mid1] == s[mid2] {
		mid1--
		mid2++
	}

	return s[mid1+1 : mid2]
}

func maxLenOf(lhs, rhs string) string {
	if len(lhs) < len(rhs) {
		return rhs
	}
	return lhs
}

func maxLenOf3(a, b, c string) string {
	return maxLenOf(a, maxLenOf(b, c))
}
