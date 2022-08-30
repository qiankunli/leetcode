package matrix

import (
	"fmt"
	"testing"
)

func TestKthSmallest(t *testing.T) {
	//matrix := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	//fmt.Println(KthSmallest(matrix, 8))

	//matrix := [][]int{{-5}}
	//fmt.Println(KthSmallest(matrix, 1))

	matrix := [][]int{{1, 91}, {90, 100}}
	fmt.Println(KthSmallest(matrix, 2))

	//matrix := [][]int{{1, 3, 5}, {6, 7, 12}, {11, 14, 14}}
	//fmt.Println(KthSmallest(matrix, 3))

}
