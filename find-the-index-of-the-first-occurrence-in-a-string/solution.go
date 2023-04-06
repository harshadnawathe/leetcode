package findtheindexofthefirstoccurrenceinastring

const PrimeRK = 16777619

func hashStr(s string) (uint32, uint32) {
	var h uint32 = 0
	for i := 0; i < len(s); i++ {
		h = h*PrimeRK + uint32(s[i])
	}

	var pow, sq uint32 = 1, PrimeRK
	for i := len(s); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}

	return h, pow
}

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	p, pow := hashStr(needle)

	h, _ := hashStr(haystack[:len(needle)])

	if h == p && haystack[:len(needle)] == needle {
		return 0
	}

	for i := len(needle); i < len(haystack); {
		h *= PrimeRK
		h += uint32(haystack[i])
		h -= pow * uint32(haystack[i-len(needle)])
		i++
		if h == p && haystack[i-len(needle):i] == needle {
			return i - len(needle)
		}
	}
	return -1
}
