package array

import (
	"fmt"
	"testing"
)

func TestFindLHS(t *testing.T) {
	//len := FindLHS([]int{1, 2, 3, 1, 2, 1})
	//len := FindLHS([]int{1, 1, 1, 1})
	len := FindLHS([]int{1, 2, 1, 3, 0, 0, 2, 2, 1, 3, 3})
	fmt.Println(len)
}
