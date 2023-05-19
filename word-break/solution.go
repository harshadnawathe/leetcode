package wordbreak

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]struct{})
	for _, word := range wordDict {
		words[word] = struct{}{}
	}

	canBreak := make([]bool, len(s)+1)
	canBreak[0] = true

	for n := 1; n <= len(s); n++ {
		for i := n - 1; i >= 0; i-- {
			if _, found := words[s[i:n]]; found && canBreak[i] {
				canBreak[n] = true
			}
		}
	}

	return canBreak[len(s)]
}

func wordBreak3(s string, wordDict []string) bool {
	words := make(map[string]struct{})
	for _, word := range wordDict {
		words[word] = struct{}{}
	}

	var f func(int) bool
	f = func(n int) bool {
		if n == 0 {
			return true
		}

		for i := n - 1; i >= 0; i-- {
			if _, found := words[s[i:n]]; found && f(i) {
				return true
			}
		}

		return false
	}

	return f(len(s))
}

func wordBreak2(s string, wordDict []string) bool {
	words := make(map[string]struct{})
	for _, word := range wordDict {
		words[word] = struct{}{}
	}

	canBreak := make([]bool, len(s)+1)
	canBreak[len(s)] = true

	for i := len(s) - 1; i >= 0; i-- {
		for n := i + 1; n <= len(s); n++ {
			if _, found := words[s[i:n]]; found && canBreak[n] {
				canBreak[i] = true
			}
		}
	}

	return canBreak[0]
}

func wordBreak1(s string, wordDict []string) bool {
	words := make(map[string]struct{})
	for _, word := range wordDict {
		words[word] = struct{}{}
	}

	var f func(int) bool
	f = func(i int) bool {
		if i == len(s) {
			return true
		}

		for n := i + 1; n <= len(s); n++ {
			if _, found := words[s[i:n]]; found && f(n) {
				return true
			}
		}

		return false
	}
	return f(0)
}
