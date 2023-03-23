package zigzagconversion

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	jump := (2 * numRows) - 2

	var b strings.Builder
	for r := 0; r < numRows; r++ {
		for c := r; c < len(s); c += jump {
			b.WriteByte(s[c])
			mid := c + (jump - 2*r)
			if mid > c && mid < c+jump && mid < len(s) {
				b.WriteByte(s[mid])
			}
		}
	}
	return b.String()
}
