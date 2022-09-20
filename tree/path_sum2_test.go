package tree

import (
	"fmt"
	"testing"
)

func TestPathSum2(t *testing.T) {
	right4 := &TreeNode{Val: 11}
	right3 := &TreeNode{Val: 1}
	right2 := &TreeNode{Val: -2}
	left3 := &TreeNode{Val: 3}
	right1 := &TreeNode{Val: 2, Right: right3}
	left1 := &TreeNode{Val: 3, Left: left3, Right: right2}
	right := &TreeNode{Val: -3, Right: right4}
	left := &TreeNode{Val: 5, Left: left1, Right: right1}
	root := &TreeNode{Val: 10, Left: left, Right: right}
	fmt.Println(PathSum2(root, 8))
}
