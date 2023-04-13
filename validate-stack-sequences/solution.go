package validatestacksequences

type stack []int

func top(stk stack) int           { return stk[len(stk)-1] }
func pop(stk stack) stack         { return stk[:len(stk)-1] }
func push(stk stack, x int) stack { return append(stk, x) }

func validateStackSequences(pushed []int, popped []int) bool {
	pu, po := 0, 0
	stk := stack{}
	for ; pu < len(pushed); pu++ {
		stk = append(stk, pushed[pu])

		for po < len(popped) && len(stk) > 0 && top(stk) == popped[po] {
			stk = pop(stk)
			po++
		}
	}
	return pu == len(pushed) && po == len(popped)
}
