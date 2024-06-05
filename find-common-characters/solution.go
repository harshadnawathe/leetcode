package findcommoncharacters

func commonChars(words []string) []string {
	acc := map[byte]int{}
	for i := range len(words[0]) {
		acc[words[0][i]] += 1
	}

	for _, word := range words {
		cc := map[byte]int{}
		for i := range len(word) {
			cc[word[i]] += 1
		}

		for c := byte('a'); c <= 'z'; c++ {
			acc[c] = min(acc[c], cc[c])
		}
	}

	result := []string{}
	for c := byte('a'); c <= 'z'; c++ {
		for count := acc[c]; count > 0; count-- {
			result = append(result, string(c))
		}
	}

	return result
}
