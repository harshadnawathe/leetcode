package longestvalidparentheses

type stack []int

func top(stk stack) int           { return stk[len(stk)-1] }
func pop(stk stack) stack         { return stk[:len(stk)-1] }
func push(stk stack, x int) stack { return append(stk, x) }

func maxOf(lhs, rhs int) int {
	if rhs < lhs {
		return lhs
	}
	return rhs
}

func longestValidParentheses(s string) int {
	max := 0
	stk := stack{}
	previousStart := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stk = push(stk, i)
			continue
		}

		if len(stk) == 0 {
			previousStart = i + 1
			continue
		}

		stk = pop(stk)
		start := previousStart
		if len(stk) > 0 {
			start = top(stk) + 1
		}

		max = maxOf(max, i-start+1)
	}

	return max
}
