package array

import (
	"fmt"
	"testing"
)

func TestSmallestRange(t *testing.T) {
	//nums := [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
	//fmt.Println(SmallestRange(nums))

	//nums := [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
	//fmt.Println(SmallestRange(nums))

	//nums := [][]int{{10, 10}, {11, 11}}
	//fmt.Println(SmallestRange(nums))

	nums := [][]int{{-5, -4, -3, -2, -1}, {1, 2, 3, 4, 5}}
	fmt.Println(SmallestRange(nums))
}
