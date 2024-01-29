package implementqueueusingstacks

type MyQueue struct {
	pushStack, popStack stack
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	q.pushStack = push(q.pushStack, x)
}

func (q *MyQueue) Pop() int {
	x := q.Peek()
	q.popStack = pop(q.popStack)
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.popStack) == 0 {
		for len(q.pushStack) > 0 {
			x := peek(q.pushStack)
			q.popStack = push(q.popStack, x)
			q.pushStack = pop(q.pushStack)
		}
	}

	x := peek(q.popStack)
	return x
}

func (q *MyQueue) Empty() bool {
	return len(q.pushStack) == 0 && len(q.popStack) == 0
}

type stack []int

func push(stk stack, x int) stack {
	return append(stk, x)
}

func pop(stk stack) stack {
	return stk[:len(stk)-1]
}

func peek(stk stack) int {
	return stk[len(stk)-1]
}
