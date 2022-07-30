package queue

type Queue struct {
	values   []int
	push     int
	pop      int
	len      int
	capacity int
}

func NewQueue(capacity int) *Queue {
	return &Queue{
		values:   make([]int, capacity+1),
		capacity: capacity,
	}
}
func (q *Queue) Push(v int) {
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
func (q *Queue) Pop() int {
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
