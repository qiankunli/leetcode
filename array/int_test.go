package array

import (
	"fmt"
	"math"
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
func TestSlice(t *testing.T) {
	array := []int{1, 2, 3, 4}

	fmt.Println(array[0:2])
	arrayCopy := make([]int, len(array))
	copy(arrayCopy, array)
	fmt.Println(arrayCopy)
}
func TestPow(t *testing.T) {
	fmt.Println(math.Pow10(-2))
}
