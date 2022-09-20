package design

type Stack struct {
	capacity int
	len      int
	array    []int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		capacity: capacity,
		array:    make([]int, capacity),
	}
}
func (s *Stack) Push(value int) {
	//todo 边界
	s.len++
	s.array[s.len] = value
}

func (s *Stack) Pop() int {
	//todo 边界
	v := s.array[s.len]
	s.len--
	return v
}
func (s *Stack) Top() int {
	return s.array[s.len]
}

type MaxStack struct {
	data *Stack
	max  *Stack
}

func NewMaxStack(capacity int) *MaxStack {
	return &MaxStack{
		data: NewStack(capacity),
		max:  NewStack(capacity),
	}
}
func (s *MaxStack) Push(value int) {
	s.data.Push(value)
	max := s.max.Top()
	if max < value {
		s.max.Push(value)
		return
	}
	s.max.Push(max)
}

func (s *MaxStack) Pop() int {
	v := s.data.Pop()
	s.max.Pop()
	return v
}
func (s *MaxStack) Max() int {
	return s.max.Top()
}
