package array

import (
	"fmt"
	"leetcode/queue"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(MaxSlidingWindow(nums, 3))
}

func TestPut(t *testing.T) {
	q := queue.NewDQueue(10)
	Put(q, 10)
	q.Print()
	Put(q, 9)
	q.Print()
	Put(q, 5)
	q.Print()
	Put(q, 7)
	q.Print()
	Put(q, 11)
	q.Print()
}
