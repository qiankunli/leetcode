package tree

import (
	"fmt"
	"testing"
)

func TestIsValidBST(t *testing.T) {
	right1 := &TreeNode{Val: 20}
	left1 := &TreeNode{Val: 6}
	right := &TreeNode{Val: 15, Left: left1, Right: right1}
	left := &TreeNode{Val: 5, Left: left1}
	root := &TreeNode{Val: 10, Left: left, Right: right}
	fmt.Println(IsValidBST(root))
}
