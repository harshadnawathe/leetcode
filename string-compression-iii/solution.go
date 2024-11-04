package stringcompressioniii

import (
	"strings"
)

func compressedString(word string) string {
	var b strings.Builder

	var begin, end int

	for ; end < len(word); end++ {
		if word[begin] == word[end] {
			continue
		}

		compressInterval(word, begin, end, &b)
		begin = end
	}

	compressInterval(word, begin, end, &b)

	return b.String()
}

func compressInterval(word string, begin, end int, b *strings.Builder) {
	n := end - begin

	for ; n >= 9; n -= 9 {
		b.WriteByte('9')
		b.WriteByte(word[begin])
	}

	if n > 0 {
		b.WriteByte(byte((n + '0')))
		b.WriteByte(word[begin])
	}
}
