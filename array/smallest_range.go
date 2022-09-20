package array

import "math"

func SmallestRange(nums [][]int) []int {
	return smallestRange(nums)
}

// 632
/**
思路比较清晰
1. 所有数组合并（增序），每一个位置除了元素，还记录了属于第几个原来的数组
2. 对新数组用滑动窗口，如果窗口里包含 所有原来数组的元素 且窗口长度最短，就是答案了
3. 就是每个步骤都不好写
*/
func smallestRange(nums [][]int) []int {
	n := len(nums)
	if n == 1 {
		return []int{nums[0][0], nums[0][0]}
	}
	head := &item{}
	numbers := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < len(nums[i]); j++ {
			appendItem(head, nums[i][j], i)
		}
		numbers = append(numbers, i)
	}
	l := head.Next
	r := head.Next
	min := math.MaxInt
	minL := l
	minR := r
	window := window{
		hash: make(map[int]int),
	}
	window.add(l.Hash)
	windowLen := 1
	for windowLen >= 1 && r != nil {
		if window.contains(numbers) {
			if min > r.Val-l.Val {
				min = r.Val - l.Val
				minL = l
				minR = r
			}
			window.delete(l.Hash)
			l = l.Next
			windowLen--
		} else {
			r = r.Next
			windowLen++
			if r != nil {
				window.add(r.Hash)
			}
		}
	}
	return []int{minL.Val, minR.Val}
}

type window struct {
	hash map[int]int
}

func (w *window) add(hash map[int]bool) {
	for k, v := range hash {
		if v == true {
			w.hash[k]++
		}
	}
}
func (w *window) delete(hash map[int]bool) {
	for k, v := range hash {
		if v == true {
			w.hash[k]--
		}
	}
}
func (w *window) contains(nums []int) bool {
	for _, v := range nums {
		if v, ok := w.hash[v]; !ok || v == 0 {
			return false
		}
	}
	return true
}

type item struct {
	Val  int
	Hash map[int]bool
	Next *item
}

func appendItem(head *item, val, number int) {

	node := &item{
		Val:  val,
		Hash: make(map[int]bool),
	}
	node.Hash[number] = true
	if head.Next == nil {
		head.Next = node
		return
	}

	cur := head.Next
	preCur := head
	for cur != nil {

		if val == cur.Val {
			cur.Hash[number] = true
			return
		}
		if val < cur.Val {
			preCur.Next = node
			node.Next = cur
			return
		}
		preCur = cur
		cur = cur.Next
	}
	preCur.Next = node
}
