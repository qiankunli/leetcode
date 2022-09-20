package tree

import (
	"fmt"
	"testing"
)

func TestWidthOfBinaryTree(t *testing.T) {
	left3 := &TreeNode{Val: 4}
	right2 := &TreeNode{Val: 7}
	left2 := &TreeNode{Val: 1}
	right1 := &TreeNode{Val: 5, Left: left3, Right: right2}
	left1 := &TreeNode{Val: 2, Left: left2}
	root := &TreeNode{Val: 3, Left: left1, Right: right1}
	fmt.Println(WidthOfBinaryTree(root))
}
