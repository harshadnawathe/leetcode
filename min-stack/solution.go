package minstack

type MinStack struct {
	values   []int
	minIndex []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (stk *MinStack) Push(val int) {
	stk.values = append(stk.values, val)
	if len(stk.minIndex) == 0 || val < stk.GetMin() {
		stk.minIndex = append(stk.minIndex, len(stk.values)-1)
	}
}

func (stk *MinStack) Pop() {
	if len(stk.values)-1 == stk.minIndex[len(stk.minIndex)-1] {
		stk.minIndex = stk.minIndex[:len(stk.minIndex)-1]
	}
	stk.values = stk.values[:len(stk.values)-1]
}

func (stk *MinStack) Top() int {
	return stk.values[len(stk.values)-1]
}

func (stk *MinStack) GetMin() int {
	return stk.values[stk.minIndex[len(stk.minIndex)-1]]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
