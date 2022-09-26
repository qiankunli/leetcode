package string

func IsValid(s string) bool {
	return isValid(s)
}

// 20  有效的括号
// 没事手撸一下
func isValid(s string) bool {

	// 用栈，碰到左就进栈，碰到右就出栈（栈顶还得一样）
	return false

}

func isValid2(s string) bool {
	if len(s) <= 1 {
		return false
	}
	stack := NewStack(len(s))
	for i := 0; i < len(s); i++ {
		if stack.Len() > 0 {
			if match(stack.Top(), s[i:i+1]) {
				stack.Pop()
				continue
			}
		}
		stack.Push(s[i : i+1])
	}
	return stack.Len() == 0
}
func match(left, right string) bool {
	if left == "(" && right == ")" {
		return true
	}
	if left == "[" && right == "]" {
		return true
	}
	if left == "{" && right == "}" {
		return true
	}
	return false
}

type Stack struct {
	array    []string
	capacity int
	top      int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		capacity: capacity,
		array:    make([]string, capacity),
	}
}
func (s *Stack) Push(ch string) {
	if s.top >= s.capacity {
		panic("stack is full")
	}
	s.array[s.top] = ch
	s.top++
}

func (s *Stack) Pop() string {
	if s.top == 0 {
		panic("stack is empty")
	}
	s.top--
	return s.array[s.top]
}
func (s *Stack) Top() string {
	return s.array[s.top-1]
}
func (s *Stack) Len() int {
	return s.top
}
