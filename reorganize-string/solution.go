package reorganizestring

func reorganizeString(s string) string {
	charCount := make([]int, 26)
	var maxChar int
	for i := 0; i < len(s); i++ {
		char := int(s[i] - 'a')
		charCount[char] += 1
		if charCount[maxChar] < charCount[char] {
			maxChar = char
		}
	}

	if charCount[maxChar] > (len(s)+1)/2 {
		return ""
	}

	bs := make([]byte, len(s))

	p := 0
	for charCount[maxChar] > 0 {
		bs[p] = byte(maxChar + 'a')
		charCount[maxChar]--
		p += 2
	}

	for char, count := range charCount {
		for ; count > 0; count-- {
			if p >= len(s) {
				p = 1
			}
			bs[p] = byte(char + 'a')
			p += 2
		}
	}

	return string(bs)
}
