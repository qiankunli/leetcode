package array

import (
	"fmt"
	"testing"
)

func TestFindKthLargest(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(FindKthLargest(nums, 2))

	//nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	//fmt.Println(FindKthLargest(nums, 4))
}
