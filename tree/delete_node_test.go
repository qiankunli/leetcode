package tree

import (
	"testing"
)

func TestDeleteNode(t *testing.T) {
	right2 := &TreeNode{Val: 7}
	left2 := &TreeNode{Val: 5}
	right1 := &TreeNode{Val: 3}
	left1 := &TreeNode{Val: 1}
	right := &TreeNode{Val: 6, Left: left2, Right: right2}
	left := &TreeNode{Val: 2, Left: left1, Right: right1}
	root := &TreeNode{Val: 4, Left: left, Right: right}
	r := DeleteNode(root, 4)
	Print(r)
}
