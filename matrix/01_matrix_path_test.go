package matrix

import (
	"fmt"
	"testing"
)

func TestMatrixPath(t *testing.T) {
	matrix := [][]int{
		{0, 0, 1},
		{0, 0, 1},
		{0, 0, 0},
	}
	path := MatrixPath(matrix)
	fmt.Println(path)
}
