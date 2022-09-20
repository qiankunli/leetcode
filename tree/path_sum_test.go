package tree

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {
	right3 := &TreeNode{Val: 1}
	right2 := &TreeNode{Val: 2}
	left3 := &TreeNode{Val: 5}
	left31 := &TreeNode{Val: 7}
	left2 := &TreeNode{Val: 13}
	right1 := &TreeNode{Val: 4, Left: left3, Right: right3}
	left1 := &TreeNode{Val: 11, Left: left31, Right: right2}
	right := &TreeNode{Val: 8, Left: left2, Right: right1}
	left := &TreeNode{Val: 4, Left: left1}
	root := &TreeNode{Val: 5, Left: left, Right: right}
	fmt.Println(PathSum(root, 22))
}
