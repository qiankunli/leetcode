package tree

import (
	"fmt"
	"testing"
)

func TestVerticalOrder(t *testing.T) {
	//right := &TreeNode{Val: 3}
	//left := &TreeNode{Val: 2}
	//root := &TreeNode{Val: 1, Left: left, Right: right}
	//fmt.Println(VerticalOrder(root))

	right1 := &TreeNode{Val: 7}
	left1 := &TreeNode{Val: 15}
	right := &TreeNode{Val: 20, Left: left1, Right: right1}
	left := &TreeNode{Val: 9}
	root := &TreeNode{Val: 3, Left: left, Right: right}
	fmt.Println(VerticalOrder(root))
}
