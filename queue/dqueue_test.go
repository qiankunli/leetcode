package queue

import (
	"fmt"
	"testing"
)

func TestDQueue_Push(t *testing.T) {
	q := NewDQueue(2)
	q.Push(10)
	q.Push(11)
	fmt.Println(q.Pop())
	q.Push(12)
}

func TestDQueue_PushLeft(t *testing.T) {
	q := NewDQueue(2)
	q.PushLeft(10)
	q.PushLeft(11)
	fmt.Println(q.PopRight())
	q.Push(12)
}

func TestDQueue_PopRight(t *testing.T) {
	q := NewDQueue(10)
	q.Push(10)
	q.PushLeft(11)
	q.PushLeft(12)
	q.Push(13)
	fmt.Println(q.PopRight())
	fmt.Println(q.Right())
	fmt.Println(q.Left())
	fmt.Println(q.values)
}
