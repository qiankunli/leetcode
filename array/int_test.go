package array

import (
	"fmt"
	"testing"
)

func TestTraverseChildSequence(t *testing.T) {
	TraverseChildSequence([]int{1, 2, 3}, func(nums []int, start, end int) {
		fmt.Println(nums[start : end+1])
	})
}

func TestTraverseChildArray(t *testing.T) {
	TraverseChildArray([]int{1, 2, 3}, func(nums []int) {
		fmt.Println(nums)
	})
}

func TestTraverseChildArrayWithCondition(t *testing.T) {
	TraverseChildArrayWithCondition([]int{1, 2, 3},
		func(nums, bits []int) bool {
			fmt.Println(bits)
			return true
		},
		func(nums []int) {
			//fmt.Println(nums)
		})
}
