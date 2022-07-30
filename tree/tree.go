package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	values   []*TreeNode
	push     int
	pop      int
	len      int
	capacity int
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		values:   make([]*TreeNode, capacity+1),
		capacity: capacity,
	}
}
func (q *Queue) Push(v *TreeNode) {
	if q.len >= q.capacity {
		panic("queue is full")
	}
	q.values[q.push] = v
	q.len++
	q.push++
	if q.push >= q.capacity {
		q.push = 0
	}
}
func (q *Queue) Pop() *TreeNode {
	if q.len <= 0 {
		panic("queue is empty")
	}
	v := q.values[q.pop]
	q.len--
	q.pop++
	if q.pop >= q.capacity {
		q.pop = 0
	}
	return v
}
func (q *Queue) Len() int {
	return q.len
}

func BFS(t *TreeNode, handler func(node *TreeNode)) {
	if t == nil {
		return
	}
	q := NewQueue(100)
	q.Push(t)
	for q.Len() > 0 {
		node := q.Pop()
		handler(node)
		if node.Left != nil {
			q.Push(node.Left)
		}
		if node.Right != nil {
			q.Push(node.Right)
		}
	}
}

func Print(t *TreeNode) {
	BFS(t, func(node *TreeNode) {
		fmt.Println(node)
	})
}
