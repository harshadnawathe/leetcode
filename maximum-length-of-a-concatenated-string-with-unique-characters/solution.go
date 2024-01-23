package maximumlengthofaconcatenatedstringwithuniquecharacters

func maxLength(arr []string) int {
	cache := make(map[[2]any]int)
	var f func(string, int) int
	f = func(s string, n int) int {
		if len(arr) == n {
			return len(s)
		}
		key := [2]any{s, n}
		if cached, ok := cache[key]; ok {
			return cached
		}

		concated := s + arr[n]
		var result int
		if containsUniqueChars(concated) {
			result = maxOf(f(concated, n+1), f(s, n+1))
		} else {
			result = f(s, n+1)
		}
		cache[key] = result
		return result
	}
	return f("", 0)
}

func containsUniqueChars(s string) bool {
	table := [26]bool{}
	for i := 0; i < len(s); i++ {
		if table[s[i]-'a'] {
			return false
		}
		table[s[i]-'a'] = true
	}
	return true
}

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}
