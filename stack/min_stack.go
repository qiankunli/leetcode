package stack

type MinStack struct {
	data *stack // 记录当前数据
	min  *stack // 记录当前最小值
}
type stack struct {
	array []int
	len   int // 指向栈顶，下一个要插入的元素
}

func (this *stack) Push(val int) {
	this.array[this.len] = val
	this.len++
}
func (this *stack) Pop() {
	this.len--
}

func (this *stack) Top() int {
	return this.array[this.len-1]
}
func (this *stack) Len() int {
	return this.len
}

func Constructor() MinStack {
	return MinStack{
		data: &stack{
			array: make([]int, 30000),
		},
		min: &stack{
			array: make([]int, 30000),
		},
	}
}

func (this *MinStack) Push(val int) {
	this.data.Push(val)
	if this.min.Len() == 0 { // 没有元素时，top 无法取值
		this.min.Push(val)
		return
	}
	top := this.min.Top()
	if val <= top { // 新元素 小于 最小栈top 值，加入最小栈
		this.min.Push(val)
	}
}

func (this *MinStack) Pop() {
	val := this.data.Top()
	this.data.Pop()
	top := this.min.Top()
	if val <= top { // 取出元素 小于等于 最小栈top 值，移除最小栈。重复了就要多次添加
		this.min.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.data.Top()
}

func (this *MinStack) GetMin() int {
	return this.min.Top()
}
