package minimumremovetomakevalidparentheses

import "strings"

// time: O(n)
// space: O(n)
func minRemoveToMakeValid(s string) string {
	stk := []int{}

	for i, char := range s {
		switch char {
		case '(':
			stk = append(stk, i)

		case ')':
			if len(stk) == 0 || s[stk[len(stk)-1]] == ')' {
				stk = append(stk, i)
			} else if s[stk[len(stk)-1]] == '(' {
				stk = stk[:len(stk)-1]
			}
		}
	}

	if len(stk) == 0 {
		return s
	}

	var b strings.Builder
	begin := 0
	for _, end := range stk {
		b.WriteString(s[begin:end])
		begin = end + 1
	}
	b.WriteString(s[begin:])

	return b.String()
}
