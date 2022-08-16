package queue

import "fmt"

type DQueue struct {
	values   []int
	push     int
	pop      int
	len      int
	capacity int
}

func NewDQueue(capacity int) *DQueue {
	return &DQueue{
		values:   make([]int, capacity),
		capacity: capacity,
	}
}

func (q *DQueue) Push(v int) {
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
func (q *DQueue) PushLeft(v int) {
	if q.len >= q.capacity {
		panic("queue is full")
	}
	// pop 指向当前可以pop 的元素，所以要先左移一位
	q.pop--
	if q.pop < 0 {
		q.pop = q.capacity + q.pop
	}
	q.values[q.pop] = v
	q.len++
}
func (q *DQueue) Pop() int {
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
func (q *DQueue) PopRight() int {
	if q.len <= 0 {
		panic("queue is empty")
	}
	// push 指向即将出入的元素，所以先左移一位
	q.push--
	if q.push < 0 {
		q.push = q.capacity + q.push
	}
	v := q.values[q.push]
	q.len--
	return v
}

func (q *DQueue) Right() int {
	if q.len <= 0 {
		panic("queue is empty")
	}
	// push 指向即将插入的元素，所以返回左边那个
	index := q.push - 1
	if index < 0 {
		index = q.capacity + index
	}
	return q.values[index]

}

func (q *DQueue) Left() int {
	if q.len <= 0 {
		panic("queue is empty")
	}
	return q.values[q.pop]
}
func (q *DQueue) Len() int {
	return q.len
}

func (q *DQueue) Print() {
	if q.len == 0 {
		fmt.Println()
	}
	result := make([]int, 0)
	if q.pop > q.push {
		for i := q.pop; i < q.capacity; i++ {
			result = append(result, q.values[i])
		}
		for i := 0; i < q.push; i++ {
			result = append(result, q.values[i])
		}
	} else {
		for i := q.pop; i < q.push; i++ {
			result = append(result, q.values[i])
		}
	}
	fmt.Println(result)
}
