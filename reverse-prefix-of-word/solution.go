package reverseprefixofword

func reversePrefix(word string, ch byte) string {
	// find
	i := 0
	for ; i < len(word) && word[i] != ch; i++ {
	}

	if i == len(word) {
		return word
	}

	// reverse
	bs := []byte(word)
	b, e := 0, i
	for b < e {
		bs[b], bs[e] = bs[e], bs[b]
		b++
		e--
	}

	return string(bs)
}
