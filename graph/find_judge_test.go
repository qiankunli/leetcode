package graph

import (
	"fmt"
	"testing"
)

func TestFindJudge(t *testing.T) {
	trust1 := [][]int{
		{1, 3},
		{2, 3},
	}
	fmt.Println(FindJudge(3, trust1))

	trust2 := [][]int{
		{1, 3},
		{2, 3},
		{3, 1},
	}
	fmt.Println(FindJudge(3, trust2))

	trust3 := [][]int{}
	fmt.Println(FindJudge(1, trust3))
}
