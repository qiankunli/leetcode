package array

import "leetcode/queue"

func MaxSlidingWindow(nums []int, k int) []int {
	return maxSlidingWindow(nums, k)
}

// 239
func maxSlidingWindow(nums []int, k int) []int {
	q := queue.NewDQueue(k)
	result := make([]int, 0)
	for i := 0; i < k; i++ {
		PutIndex(q, nums, i)
	}
	result = append(result, nums[q.Left()])
	for i := k; i < len(nums); i++ {
		// 准备添加右侧元素了
		if i-k == q.Left() {
			// 若最左侧元素 在窗口之外，则移出
			q.Pop()
		}
		// 在最右侧加入新元素
		PutIndex(q, nums, i)
		// 添加当前窗口的最大元素
		result = append(result, nums[q.Left()])
	}
	return result
}

func Put(q *queue.DQueue, value int) {
	if q.Len() == 0 {
		q.Push(value)
		return
	}
	for q.Len() > 0 {
		if value <= q.Right() {
			q.Push(value)
			return
		}
		q.PopRight()
	}
	q.Push(value)
}

func PutIndex(q *queue.DQueue, nums []int, index int) {
	if q.Len() == 0 {
		q.Push(index)
		return
	}
	value := nums[index]
	for q.Len() > 0 {
		if value <= nums[q.Right()] {
			q.Push(index)
			return
		}
		q.PopRight()
	}
	q.Push(index)
}
