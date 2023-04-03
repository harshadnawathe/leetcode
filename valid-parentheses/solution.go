package validparentheses

func isValid(s string) bool {
	var stk stack
	for i := 0; i < len(s); i++ {
		if _, isOpening := parentheses[s[i]]; isOpening {
			stk = push(stk, s[i])
			continue
		}

		if len(stk) == 0 {
			return false
		}

		lastOpening := top(stk)
		stk = pop(stk)
		if s[i] != parentheses[lastOpening] {
			return false
		}
	}
	return len(stk) == 0
}

var parentheses = map[byte]byte{
	'(': ')',
	'[': ']',
	'{': '}',
}

type stack []byte

func top(stk stack) byte {
	return stk[len(stk)-1]
}

func push(stk stack, x byte) stack {
	return append(stk, x)
}

func pop(stk stack) stack {
	return stk[:len(stk)-1]
}
