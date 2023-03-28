package regularexpressionmatching

func isMatch(s string, p string) bool {
	matchResult := make([][]bool, len(s)+1)
	for i := range matchResult {
		matchResult[i] = make([]bool, len(p)+1)
	}

	matchResult[len(s)][len(p)] = true

	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			firstMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')
			if j+1 < len(p) && p[j+1] == '*' {
				matchResult[i][j] = matchResult[i][j+2] || (firstMatch && matchResult[i+1][j])
			} else {
				matchResult[i][j] = firstMatch && matchResult[i+1][j+1]
			}
		}
	}

	return matchResult[0][0]
}

func isMatch1(s string, p string) bool {
	var f func(int, int) bool
	f = func(i, j int) bool {
		if j == len(p) {
			return i == len(s)
		}

		firstMatch := i < len(s) && (s[i] == p[j] || p[j] == '.')

		if j+1 < len(p) && p[j+1] == '*' {
			return f(i, j+2) || (firstMatch && f(i+1, j))
		}
		return firstMatch && f(i+1, j+1)
	}
	return f(0, 0)
}

func isMatch2(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	firstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) > 1 && p[1] == '*' {
		return isMatch2(s, p[2:]) || (firstMatch && isMatch2(s[1:], p))
	} else {
		return firstMatch && isMatch2(s[1:], p[1:])
	}
}
