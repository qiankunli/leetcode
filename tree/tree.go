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

func DFS(t *TreeNode, handler func(node *TreeNode)) {
	if t == nil {
		return
	}
	handler(t)
	DFS(t.Left, handler)
	DFS(t.Right, handler)
}

func DFS2(t *TreeNode, nums []int, predict func(values []int) bool) {
	if t == nil {
		return
	}
	nums = append(nums, t.Val)
	if predict(nums) { // 满足条件，结束
		return
	}
	DFS2(t.Left, nums, predict)
	DFS2(t.Right, nums, predict)
	nums = nums[:len(nums)-1]
}

func Print(t *TreeNode) {
	BFS(t, func(node *TreeNode) {
		fmt.Println(node)
	})
}

func NewTree(array []int, nullValue int) *TreeNode {
	root := &TreeNode{}
	i := 0
	BFS2(root, func() (int, int, int, int) {
		v := array[i]
		index := i
		i++
		return index, v, len(array), nullValue
	})
	return root
}

// BFS 的反向应用
func BFS2(t *TreeNode, provider func() (int, int, int, int)) {
	if t == nil {
		return
	}
	q := NewQueue(100)
	q.Push(t)
	for q.Len() > 0 {
		node := q.Pop()
		index, value, len, nullValue := provider()
		node.Val = value
		if value == nullValue {
			continue
		}
		// 每创建一个node 都要确保有val 可以赋值。当前元素的left 和 right 加到 queue 里就需要 2个val，queue 里的每个元素也需要一个val
		qLen := q.Len()
		if qLen+index <= len-2 {
			node.Left = &TreeNode{}
			q.Push(node.Left)
		}
		if qLen+index <= len-3 {
			node.Right = &TreeNode{}
			q.Push(node.Right)
		}
	}
}
func Count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + Count(root.Right)
	}
	if root.Right == nil {
		return 1 + Count(root.Left)
	}
	return 1 + Count(root.Left) + Count(root.Right)
}
func Count2(root *TreeNode, nullValue int) int {
	if root == nil || root.Val == nullValue {
		return 0
	}
	if root.Left == nil || root.Left.Val == nullValue {
		return 1 + Count2(root.Right, nullValue)
	}
	if root.Right == nil || root.Right.Val == nullValue {
		return 1 + Count2(root.Left, nullValue)
	}
	return 1 + Count2(root.Left, nullValue) + Count2(root.Right, nullValue)
}

func Traverse(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	path := make([]int, 0)
	result := make([][]int, 0)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil { // 处理逻辑如果加在这里，会被处理两次
			return
		}
		path = append(path, root.Val)
		// 加逻辑的地方
		if root.Left == nil && root.Right == nil {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
		}
		dfs(root.Left)
		dfs(root.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return result
}
