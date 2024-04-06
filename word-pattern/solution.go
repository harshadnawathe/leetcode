package wordpattern

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	letterToWord := make(map[rune]string)
	wordToLetter := make(map[string]rune)

	for i, char := range pattern {
		letter, letterOk := wordToLetter[words[i]]
		word, wordOk := letterToWord[char]

		if !letterOk && !wordOk {
			letterToWord[char] = words[i]
			wordToLetter[words[i]] = char
			continue
		}

		if letter != char || word != words[i] {
			return false
		}

	}

	return true
}
