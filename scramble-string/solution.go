package scramblestring

func isScramble(s1 string, s2 string) bool {
	cache := make(cache)

	var f func(string, string) bool
	f = func(s1, s2 string) bool {
		k := key{s1, s2}

		if cached, found := cache[k]; found {
			return cached
		}
		if s1 == s2 {
			cache[k] = true
			return cache[k]
		}

		if !isPermutation(s1, s2) {
			cache[k] = false
			return cache[k]
		}
		for i := 1; i < len(s1); i++ {
			s1First, s1Second := s1[:i], s1[i:]
			if (f(s1First, s2[:i]) && f(s1Second, s2[i:])) || (f(s1First, s2[len(s2)-i:]) && f(s1Second, s2[:len(s2)-i])) {
				return true
			}

		}
		cache[k] = false
		return cache[k]
	}

	return f(s1, s2)
}

// cache
type cache map[key]bool

type key struct {
	s1, s2 string
}

// isPermutation
func isPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var m bitmap
	for i := 0; i < len(s1); i++ {
		flip(m, uint(s1[i]-'a'))
		flip(m, uint(s2[i]-'a'))
	}

	return m == 0
}

// bitmap
type bitmap uint

func set(m bitmap, i uint) bitmap {
	x := uint(m)
	x = x | (1 << i)
	return bitmap(x)
}

func unset(m bitmap, i uint) bitmap {
	x := uint(m)
	x = x & (1 << i)
	return bitmap(x)
}

func isSet(m bitmap, i uint) bool {
	return (uint(m) & (1 << i)) != 0
}

func flip(m bitmap, i uint) bitmap {
	if isSet(m, i) {
		return unset(m, i)
	}
	return set(m, i)
}
