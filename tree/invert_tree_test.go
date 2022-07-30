package tree

import "testing"

func TestInvertTree(t *testing.T) {
	right2 := &TreeNode{Val: 3}
	left2 := &TreeNode{Val: 2}
	right1 := &TreeNode{Val: 3}
	left1 := &TreeNode{Val: 2, Left: left2, Right: right2}
	root := &TreeNode{Val: 1, Left: left1, Right: right1}
	InvertTree(root)
}
