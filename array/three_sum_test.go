package array

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	result := ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	//result := ThreeSum([]int{-2, 0, 0, 2, 2})
	fmt.Println(result)
}
