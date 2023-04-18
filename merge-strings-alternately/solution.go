package mergestringsalternately

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var b strings.Builder

	var i, j int
	for i < len(word1) && j < len(word2) {
		b.WriteByte(word1[i])
		b.WriteByte(word2[j])
		i++
		j++
	}

	b.WriteString(word1[i:])
	b.WriteString(word2[j:])

	return b.String()
}
