package string

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	fmt.Println(IsValid("()"))
	fmt.Println(IsValid("()[]{}"))
	fmt.Println(IsValid("(]"))
}
func TestStack_Push(t *testing.T) {
	stack := NewStack(10)
	stack.Push("a")
	fmt.Println(stack.Pop())
	stack.Push("b")
	stack.Push("c")
	fmt.Println(stack.Top())
	fmt.Println(stack.Len())
	fmt.Println(stack.Pop())
}
