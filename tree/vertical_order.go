package tree

import (
	"math"
)

func VerticalOrder(root *TreeNode) [][]int {
	return verticalOrder(root)
}

// 314
// BFS
func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make(map[int][]int)
	min := math.MaxInt // 维护位置最小值
	max := math.MinInt // 维护位置最大值
	queue := NewTreeLocationNodeQueue(100)
	queue.Push(&TreeLocationNode{TreeNode: root, Location: 0})
	for queue.len > 0 {
		node := queue.Pop()
		location := node.Location
		if min > location {
			min = location
		}
		if max < location {
			max = location
		}
		if _, ok := result[location]; !ok {
			result[location] = make([]int, 0)
		}
		result[location] = append(result[location], node.Val)
		if node.Left != nil {
			queue.Push(&TreeLocationNode{TreeNode: node.Left, Location: node.Location - 1})
		}
		if node.Right != nil {
			queue.Push(&TreeLocationNode{TreeNode: node.Right, Location: node.Location + 1})
		}
	}

	// 将[min,max] 转换为 [0,len(result)]
	data := make([][]int, len(result))
	c := 0
	for i := min; i <= max; i++ {
		if _, ok := result[i]; ok {
			data[c] = result[i]
			c++
		}
	}
	return data
}

type TreeLocationNode struct {
	*TreeNode
	Location int
}
type TreeLocationNodeQueue struct {
	values   []*TreeLocationNode
	push     int
	pop      int
	len      int
	capacity int
}

func NewTreeLocationNodeQueue(capacity int) *TreeLocationNodeQueue {
	return &TreeLocationNodeQueue{
		values:   make([]*TreeLocationNode, capacity),
		capacity: capacity,
	}
}
func (q *TreeLocationNodeQueue) Push(v *TreeLocationNode) {
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
func (q *TreeLocationNodeQueue) Pop() *TreeLocationNode {
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

// DFS 可以做但无法实现有序的结果
func verticalOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make(map[int][]int)
	min := math.MaxInt // 维护位置最小值
	max := math.MinInt // 维护位置最大值
	var doVerticalOrder func(root *TreeNode, location int, result map[int][]int)
	doVerticalOrder = func(root *TreeNode, location int, result map[int][]int) {
		if root == nil {
			return
		}
		if min > location {
			min = location
		}
		if max < location {
			max = location
		}
		if _, ok := result[location]; !ok {
			result[location] = make([]int, 0)
		}
		result[location] = append(result[location], root.Val)
		doVerticalOrder(root.Left, location-1, result)
		doVerticalOrder(root.Right, location+1, result)
	}
	doVerticalOrder(root, 0, result)
	// 将[min,max] 转换为 [0,len(result)]
	data := make([][]int, len(result))
	c := 0
	for i := min; i <= max; i++ {
		if _, ok := result[i]; ok {
			data[c] = result[i]
			c++
		}
	}
	return data
}
