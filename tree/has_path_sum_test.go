package tree

import (
	"fmt"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	//root := NewTree([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1}, -1)
	//Print(root)
	//fmt.Println(HasPathSum(root, 22))
	root2 := NewTree([]int{1, 2}, -1)
	fmt.Println(HasPathSum(root2, 1))
}
