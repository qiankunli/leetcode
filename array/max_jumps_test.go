package array

import (
	"fmt"
	"testing"
)

func TestMaxJumps(t *testing.T) {
	//nums := []int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}
	//fmt.Println(MaxJumps(nums, 2))

	nums := []int{3, 3, 3, 3, 3}
	fmt.Println(MaxJumps(nums, 3))
}
