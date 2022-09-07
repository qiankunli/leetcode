package array

import (
	"fmt"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	//grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	//fmt.Println(SpiralOrder(grid))

	grid := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(SpiralOrder(grid))

	//grid := [][]int{{7}, {9}, {6}}
	//fmt.Println(SpiralOrder(grid))
}
