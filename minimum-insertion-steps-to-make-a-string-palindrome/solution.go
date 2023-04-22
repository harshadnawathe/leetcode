package minimuminsertionstepstomakeastringpalindrome

func minInsertions(s string) int {
	minInsert := make([][]int, len(s)+1)
	for i := range minInsert {
		minInsert[i] = make([]int, len(s)+1)
	}

	for begin := len(s) - 1; begin >= 0; begin-- {
		for end := 0; end <= len(s); end++ {
			if end <= begin || end-begin == 1 {
				minInsert[begin][end] = 0
				continue
			}

			if s[begin] == s[end-1] {
				minInsert[begin][end] = minInsert[begin+1][end-1]
				continue
			}

			minInsert[begin][end] = 1 + minOf(
				minInsert[begin+1][end],
				minInsert[begin][end-1],
			)
		}
	}

	return minInsert[0][len(s)]
}

func minInsertions1(s string) int {
	var minInsert func(int, int) int
	minInsert = func(begin, end int) int {
		if end <= begin || end-begin == 1 {
			return 0
		}

		if s[begin] == s[end-1] {
			return minInsert(begin+1, end-1)
		}

		return 1 + minOf(
			minInsert(begin+1, end),
			minInsert(begin, end-1),
		)
	}
	return minInsert(0, len(s))
}

func minOf(lhs, rhs int) int {
	if rhs < lhs {
		return rhs
	}
	return lhs
}
