package tree

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	right2 := &TreeNode{Val: 2}
	right1 := &TreeNode{Val: 4}
	left1 := &TreeNode{Val: 1, Right: right2}
	root := &TreeNode{Val: 3, Left: left1, Right: right1}
	fmt.Println(KthSmallest(root, 1))
}
