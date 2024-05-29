package removeoutermostparentheses

import "strings"

func removeOuterParentheses(s string) string {
	outerParenPos := -1
	nestingCount := 0

	var b strings.Builder

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			nestingCount++
			if nestingCount == 1 {
				outerParenPos = i
			}
		} else {
			nestingCount--
			if nestingCount == 0 {
				b.WriteString(s[outerParenPos+1 : i])
			}
		}
	}

	return b.String()
}
