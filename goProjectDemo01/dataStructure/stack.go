package datastructure

type MyStack struct {
	nums []int
	size int
}

func NewMyStack() *MyStack {
	/* 	stack := new(MyStack)
	   	stack.nums = make([]int, 8)
	   	stack.size = 0 */
	//
	return &MyStack{
		nums: make([]int, 0),
		size: 0,
	}
}

func (stack *MyStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *MyStack) Size() int {
	return stack.size
}

func (stack *MyStack) Top() int {
	if stack.IsEmpty() {
		return -1
	}
	return stack.nums[stack.size-1]
}

func (stack *MyStack) Push(num int) {
	stack.nums = append(stack.nums, num)
	stack.size++
}

func (stack *MyStack) Pop() int {
	if stack.IsEmpty() {
		return -1
	}
	res := stack.Top()
	stack.nums = stack.nums[:stack.size-1]
	stack.size--
	return res
}
