package array

import (
	"fmt"
	"testing"
)

func TestRotate(t *testing.T) {
	//nums := []int{1, 2, 3, 4, 5, 6, 7}
	//Rotate(nums, 3)
	//fmt.Println(nums)

	//nums := []int{-1, -100, 3, 99}
	//Rotate(nums, 2)
	//fmt.Println(nums)

	//nums := []int{1}
	//Rotate(nums, 0)
	//fmt.Println(nums)

	//nums := []int{1, 2, 3, 4, 5, 6}
	//Rotate(nums, 4)
	//fmt.Println(nums)

	//nums := []int{1}
	//Rotate(nums, 1)
	//fmt.Println(nums)

	nums := []int{1, 2}
	Rotate(nums, 2)
	fmt.Println(nums)

}
