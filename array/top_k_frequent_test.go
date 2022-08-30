package array

import (
	"fmt"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(TopKFrequent(nums, 2))

	//nums := []int{1}
	//fmt.Println(TopKFrequent(nums, 1))
}
