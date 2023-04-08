package substringwithconcatenationofallwords

func findSubstring(s string, words []string) []int {
	wordLen := len(words[0])
	wordCnt := len(words)
	ssLen := wordLen * wordCnt

	result := []int{}
	for i := 0; i < len(s)-ssLen+1; i++ {
		table := wordsTable(words)
		cnt := 0
		for j := i; j < i+ssLen; j += wordLen {
			w := s[j : j+wordLen]
			if table[w] == 0 {
				break
			}
			table[w] -= 1
			cnt++
		}
		if cnt == wordCnt {
			result = append(result, i)
		}
	}
	return result
}

func wordsTable(words []string) map[string]int {
	m := make(map[string]int)
	for _, word := range words {
		m[word] += 1
	}
	return m
}
