package lettercombinationsofaphonenumber

var letters = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	combinations := []string{""}

	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		tmp := []string{}
		for _, letter := range letters[digit] {
			for _, combination := range combinations {
				tmp = append(tmp, combination+letter)
			}
		}
		combinations = tmp
	}

	return combinations
}

// Backtracking
func letterCombinations1(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	combinations := []string{}

	var f func(string)
	f = func(s string) {
		if len(s) == len(digits) {
			combinations = append(combinations, s)
			return
		}

		for _, letter := range letters[digits[len(s)]] {
			f(s + letter)
		}
	}
	f("")
	return combinations
}
