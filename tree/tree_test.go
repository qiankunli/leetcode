package tree

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	right2 := &TreeNode{Val: 5}
	left2 := &TreeNode{Val: 4}
	right1 := &TreeNode{Val: 3}
	left1 := &TreeNode{Val: 2, Left: left2, Right: right2}
	root := &TreeNode{Val: 1, Left: left1, Right: right1}
	BFS(root, func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}

func TestDFS(t *testing.T) {
	right2 := &TreeNode{Val: 5}
	left2 := &TreeNode{Val: 4}
	right1 := &TreeNode{Val: 3}
	left1 := &TreeNode{Val: 2, Left: left2, Right: right2}
	root := &TreeNode{Val: 1, Left: left1, Right: right1}
	DFS(root, func(node *TreeNode) {
		fmt.Println(node.Val)
	})
}

func TestDFS2(t *testing.T) {
	right2 := &TreeNode{Val: 5}
	left2 := &TreeNode{Val: 4}
	right1 := &TreeNode{Val: 3}
	left1 := &TreeNode{Val: 2, Left: left2, Right: right2}
	root := &TreeNode{Val: 1, Left: left1, Right: right1}
	nums := make([]int, 0)
	DFS2(root, nums, func(values []int) bool {
		fmt.Println(values)
		return false
	})
}

func TestNewTree(t *testing.T) {
	//root := NewTree([]int{1, -1, 0, 0, 1}, -1)
	//root := NewTree([]int{1, 0, 1, 0, 0, 0, 1}, -1)
	//root := NewTree([]int{1, 1, 0, 1, 1, 0, 1, 0}, -1)
	root := NewTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1}, -1)
	Print(root)

}

func TestCount2(t *testing.T) {
	root := NewTree([]int{3, 1, 4, -1, 2}, -1)
	fmt.Println(Count2(root, -1))
}
