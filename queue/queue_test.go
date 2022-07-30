package queue

import (
	"fmt"
	"testing"
)

func TestQueue_Push(t *testing.T) {
	q := NewQueue(2)
	q.Push(10)
	q.Push(11)
	fmt.Println(q.Pop())
	q.Push(12)
}
