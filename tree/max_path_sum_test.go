package tree

import (
	"fmt"
	"testing"
)

func TestMaxPathSum(t *testing.T) {
	//right := &TreeNode{Val: 3}
	//left := &TreeNode{Val: 2}
	//root := &TreeNode{Val: 1, Left: left, Right: right}
	//fmt.Println(MaxPathSum(root))

	//right1 := &TreeNode{Val: 7}
	//left1 := &TreeNode{Val: 15}
	//right := &TreeNode{Val: 20, Left: left1, Right: right1}
	//left := &TreeNode{Val: 9}
	//root := &TreeNode{Val: -10, Left: left, Right: right}
	//fmt.Println(MaxPathSum(root))

	//left := &TreeNode{Val: -1}
	//root := &TreeNode{Val: -2, Left: left}
	//fmt.Println(MaxPathSum(root))

	//right := &TreeNode{Val: 3}
	//left := &TreeNode{Val: -2}
	//root := &TreeNode{Val: 1, Left: left, Right: right}
	//fmt.Println(MaxPathSum(root))

	//left3 := &TreeNode{Val: -1}
	//left2 := &TreeNode{Val: -2}
	//right1 := &TreeNode{Val: 3}
	//left1 := &TreeNode{Val: 1, Left: left3}
	//right := &TreeNode{Val: -3, Left: left2}
	//left := &TreeNode{Val: -2, Left: left1, Right: right1}
	//root := &TreeNode{Val: 1, Left: left, Right: right}
	//fmt.Println(MaxPathSum(root))

	right4 := &TreeNode{Val: 1}
	right3 := &TreeNode{Val: 2}
	left3 := &TreeNode{Val: 7}
	right2 := &TreeNode{Val: 4, Right: right4}
	left2 := &TreeNode{Val: 13}
	left1 := &TreeNode{Val: 11, Left: left3, Right: right3}
	right := &TreeNode{Val: 8, Left: left2, Right: right2}
	left := &TreeNode{Val: 4, Left: left1}
	root := &TreeNode{Val: 5, Left: left, Right: right}
	fmt.Println(MaxPathSum(root))
}
