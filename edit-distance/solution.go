package editdistance

func minDistance(word1 string, word2 string) int {
	md := make([][]int, len(word1)+1)
	for r := range md {
		md[r] = make([]int, len(word2)+1)
	}

	md[0][0] = 0

	for c := 1; c < len(md[0]); c++ {
		md[0][c] = c
	}

	for r := 1; r < len(md); r++ {
		md[r][0] = r
	}

	for i := 1; i < len(md); i++ {
		for j := 1; j < len(md[0]); j++ {
			if word1[i-1] == word2[j-1] {
				md[i][j] = md[i-1][j-1]
			} else {
				md[i][j] = 1 + min(
					md[i][j-1],
					md[i-1][j],
					md[i-1][j-1],
				)
			}
		}
	}

	return md[len(word1)][len(word2)]
}
