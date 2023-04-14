package longestpalindromicsubsequence

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}

func longestPalindromeSubseq(s string) int {
	maxLen := make([][]int, len(s)+1)
	for i := range maxLen {
		maxLen[i] = make([]int, len(s)+1)
	}

	for last := 1; last <= len(s); last++ {
		for first := last; first >= 0; first-- {
			if size := last - first; size <= 1 {
				maxLen[first][last] = size
				continue
			}

			if s[first] == s[last-1] {
				maxLen[first][last] = 2 + maxLen[first+1][last-1]
				continue
			}

			maxLen[first][last] = maxOf(
				maxLen[first+1][last],
				maxLen[first][last-1],
			)
		}
	}
	return maxLen[0][len(s)]
}

// recursion with closure
func longestPalindromeSubseq2(s string) int {
	var maxLen func(int, int) int
	maxLen = func(first, last int) int {
		if size := last - first; size <= 1 {
			return size
		}

		if s[first] == s[last-1] {
			return 2 + maxLen(first+1, last-1)
		}

		return maxOf(
			maxLen(first+1, last),
			maxLen(first, last-1),
		)
	}
	return maxLen(0, len(s))
}

// recursion
func longestPalindromeSubseq1(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	if s[0] == s[len(s)-1] {
		return 2 + longestPalindromeSubseq1(s[1:len(s)-1])
	}

	return maxOf(
		longestPalindromeSubseq1(s[1:]),
		longestPalindromeSubseq1(s[:len(s)-1]),
	)
}
