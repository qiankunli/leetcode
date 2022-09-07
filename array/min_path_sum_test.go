package array

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	//grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	//fmt.Println(MinPathSum(grid))

	grid := [][]int{{1, 3, 6}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(MinPathSum(grid))
}
