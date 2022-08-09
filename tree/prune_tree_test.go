package tree

import "testing"

func TestPruneTree(t *testing.T) {
	left4 := &TreeNode{Val: 1}
	left3 := &TreeNode{Val: 0, Right: left4}
	right2 := &TreeNode{Val: 5}
	left2 := &TreeNode{Val: 4}
	right1 := &TreeNode{Val: 3, Left: left3}
	left1 := &TreeNode{Val: 2, Left: left2, Right: right2}
	root := &TreeNode{Val: 1, Left: left1, Right: right1}
	Print(root)

	root = PruneTree(root)
	Print(root)
}
