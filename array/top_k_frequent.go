package array

import "container/heap"

// 347
func TopKFrequent(nums []int, k int) []int {
	return topKFrequent(nums, k)
}
func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	if k == 1 && len(nums) == 1 {
		return nums
	}
	count := make(map[int]int)
	for _, v := range nums {
		count[v]++
	}
	h := &IHeap{}
	heap.Init(h)

	for key, value := range count {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// push 到尾部
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

// 从尾部取
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
