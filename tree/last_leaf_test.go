package tree

import (
	"fmt"
	"testing"
)

func TestLastLeaf(t *testing.T) {
	//left2 := &TreeNode{Val: 2}
	//right1 := &TreeNode{Val: 4}
	//left1 := &TreeNode{Val: 1, Left: left2}
	//root := &TreeNode{Val: 3, Left: left1, Right: right1}
	//fmt.Println(LastLeaf(root))

	left3 := &TreeNode{Val: 6}
	right2 := &TreeNode{Val: 5}
	left2 := &TreeNode{Val: 2}
	right1 := &TreeNode{Val: 4, Left: left3}
	left1 := &TreeNode{Val: 1, Left: left2, Right: right2}
	root := &TreeNode{Val: 3, Left: left1, Right: right1}
	fmt.Println(LastLeaf(root))
}
