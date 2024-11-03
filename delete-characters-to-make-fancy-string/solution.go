package deletecharacterstomakefancystring

import "strings"

func makeFancyString(s string) string {
	var b strings.Builder

	begin, end := 0, 0
	for ; end < len(s); end++ {
		if s[begin] != s[end] {
			if end-begin+1 > 2 {
				b.WriteString(s[begin : begin+2])
			} else {
				b.WriteString(s[begin:end])
			}
			begin = end
		}
	}

	if end-begin+1 > 2 {
		b.WriteString(s[begin : begin+2])
	} else {
		b.WriteString(s[begin:end])
	}

	return b.String()
}
