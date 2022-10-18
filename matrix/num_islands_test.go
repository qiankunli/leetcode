package matrix

import (
	"fmt"
	"leetcode/graph"
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	fmt.Println(graph.NumIslands(grid))
}
