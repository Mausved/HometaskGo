package stack

type Stack []interface{}

func (stack Stack) IsEmpty() bool {
	return stack.size() == 0
}

func (stack *Stack) Push(val interface{}) {
	*stack = append(*stack, val)
}

func (stack Stack) size() int {
	return len(stack)
}

func (stack *Stack) Pop() {
	if !stack.IsEmpty() {
		*stack = (*stack)[:stack.size()-1]
	}
}

func (stack Stack) GetTop() interface{} {
	return stack[stack.size()-1]
}

func (stack *Stack) ToSlice() []interface{} {
	var slice []interface{}
	for !stack.IsEmpty() {
		slice = append(slice, stack.GetTop())
		stack.Pop()
	}
	return slice
}
