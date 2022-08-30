package array

import (
	"fmt"
	"testing"
)

func TestFindClosestElements(t *testing.T) {
	nums := []int{3, 5, 8, 10}
	fmt.Println(FindClosestElements(nums, 2, 15))

	//nums := []int{1, 2, 3, 4, 5}
	//fmt.Println(FindClosestElements(nums, 4, 3))

	//nums := []int{1, 2, 3, 4, 5}
	//fmt.Println(FindClosestElements(nums, 4, -1))

	//nums := []int{-2, -1, 1, 2, 3, 4, 5}
	//fmt.Println(FindClosestElements(nums, 7, 3))

	//nums := []int{0, 0, 1, 2, 3, 3, 4, 7, 7, 8}
	//fmt.Println(FindClosestElements(nums, 3, 5))
}
