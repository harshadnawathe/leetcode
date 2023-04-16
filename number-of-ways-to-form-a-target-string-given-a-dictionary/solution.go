package numberofwaystoformatargetstringgivenadictionary

const Mod = 1000000007

// tabular dp
func numWays(words []string, target string) int {
	charCount := make([]map[byte]int, len(words[0]))
	for i := range charCount {
		charCount[i] = make(map[byte]int)
	}

	for c := 0; c < len(words[0]); c++ {
		for r := range words {
			charCount[c][words[r][c]] += 1
		}
	}

	nWays := make([][]int, len(target)+1)
	for i := range nWays {
		nWays[i] = make([]int, len(words[0])+1)
	}

	for c := 0; c <= len(words[0]); c++ {
		nWays[len(target)][c] = 1
	}

	for i := len(target) - 1; i >= 0; i-- {
		for c := len(words[0]) - 1; c >= 0; c-- {
			nWays[i][c] = nWays[i][c+1]
			nWays[i][c] += charCount[c][target[i]] * nWays[i+1][c+1]
			nWays[i][c] %= Mod
		}
	}

	return nWays[0][0]
}

// recursion + remove loop with charCount
func numWays1(words []string, target string) int {
	charCount := make([]map[byte]int, len(words[0]))
	for i := range charCount {
		charCount[i] = make(map[byte]int)
	}

	for c := 0; c < len(words[0]); c++ {
		for r := range words {
			charCount[c][words[r][c]] += 1
		}
	}

	var nWays func(int, int) int
	nWays = func(i, c int) int {
		if i == len(target) {
			return 1
		}

		if c == len(words[0]) {
			if i == len(target) {
				return 1
			} else {
				return 0
			}
		}

		n := nWays(i, c+1)
		n += charCount[c][target[i]] * nWays(i+1, c+1)
		n %= Mod

		return n
	}

	return nWays(0, 0)
}

// recursion
func numWays2(words []string, target string) int {
	var nWays func(int, int) int
	nWays = func(i, c int) int {
		if i == len(target) {
			return 1
		}

		if c == len(words[0]) {
			if i == len(target) {
				return 1
			} else {
				return 0
			}
		}

		n := nWays(i, c+1)
		for r := 0; r < len(words); r++ {
			if target[i] == words[r][c] {
				n += nWays(i+1, c+1)
				n %= Mod
			}
		}
		return n
	}

	return nWays(0, 0)
}
