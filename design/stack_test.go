package design

import (
	"fmt"
	"testing"
)

func TestMaxStack_Push(t *testing.T) {
	maxStack := NewMaxStack(100)
	maxStack.Push(100)
	maxStack.Push(10)
	maxStack.Push(200)
	maxStack.Pop()
	fmt.Println(maxStack.Max())
	maxStack.Pop()
	fmt.Println(maxStack.Max())
}
