package tree

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	right4 := &TreeNode{Val: 1}
	right3 := &TreeNode{Val: 2}
	left3 := &TreeNode{Val: 7}
	right2 := &TreeNode{Val: 4, Right: right4}
	left2 := &TreeNode{Val: 13}
	left1 := &TreeNode{Val: 11, Left: left3, Right: right3}
	right := &TreeNode{Val: 8, Left: left2, Right: right2}
	left := &TreeNode{Val: 4, Left: left1}
	root := &TreeNode{Val: 5, Left: left, Right: right}
	fmt.Println(LevelOrder(root))
}
