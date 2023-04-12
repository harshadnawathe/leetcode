package countandsay

import (
	"fmt"
	"strings"
)

func countAndSay(n int) string {
	s := "1"
	for i := 1; i < n; i++ {
		s = say(s)
	}
	return s
}

func say(ns string) string {
	var b strings.Builder

	for i := 0; i < len(ns); {
		j := i + 1
		for j < len(ns) && ns[i] == ns[j] {
			j++
		}

		b.WriteString(fmt.Sprintf("%d%c", j-i, ns[i]))

		i = j
	}
	return b.String()
}
