package removingstarsfromastring

type stack []byte

func push(stk stack, x byte) stack { return append(stk, x) }
func pop(stk stack) stack          { return stk[:len(stk)-1] }

func removeStars(s string) string {
	stk := stack{}
	for i := 0; i < len(s); i++ {
		if s[i] != '*' {
			stk = push(stk, s[i])
			continue
		}

		stk = pop(stk)
	}
	return string(stk)
}
