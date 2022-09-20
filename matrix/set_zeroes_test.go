package matrix

import (
	"fmt"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	//matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	//SetZeroes(matrix)
	//fmt.Println(matrix)

	matrix := [][]int{{1, 1, 0}, {1, 0, 1}, {1, 1, 1}}
	SetZeroes(matrix)
	fmt.Println(matrix)
}
