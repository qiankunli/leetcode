package matrix

import (
	"fmt"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	matrix := [][]int{{1, 0}}
	fmt.Println(UniquePathsWithObstacles(matrix))
}
