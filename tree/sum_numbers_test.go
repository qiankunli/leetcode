package tree

import (
	"fmt"
	"testing"
)

func TestSumNumbers(t *testing.T) {

	right2 := &TreeNode{Val: 1}
	left2 := &TreeNode{Val: 5}
	right1 := &TreeNode{Val: 0}
	left1 := &TreeNode{Val: 9, Left: left2, Right: right2}
	root := &TreeNode{Val: 4, Left: left1, Right: right1}

	fmt.Println(SumNumbers(root))

}
