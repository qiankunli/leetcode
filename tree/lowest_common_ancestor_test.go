package tree

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {

	right2 := &TreeNode{Val: 5}
	left2 := &TreeNode{Val: 4}
	right1 := &TreeNode{Val: 3}
	left1 := &TreeNode{Val: 2, Left: left2, Right: right2}
	root := &TreeNode{Val: 1, Left: left1, Right: right1}
	fmt.Println(LowestCommonAncestor(root, left1, left2))

	//root := NewTree([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}, -1)
	//fmt.Println(LowestCommonAncestor(root, root.Left, root.Left.Right.Right))

	//root := NewTree([]int{37, -34, -48, -1, -100, -101, 48, -1, -1, -1, -1, -54, -1, -71, -22, -1, -1, -1, 8}, -1)
	//fmt.Println(LowestCommonAncestor(root, root.Right.Right.Left.Left, root.Right.Right))

}
