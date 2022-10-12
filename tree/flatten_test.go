package tree

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	right2 := &TreeNode{Val: 6}
	right1 := &TreeNode{Val: 4}
	left1 := &TreeNode{Val: 3}
	right := &TreeNode{Val: 5, Right: right2}
	left := &TreeNode{Val: 2, Left: left1, Right: right1}
	root := &TreeNode{Val: 1, Left: left, Right: right}
	Flatten(root)
	Print(root)
}
