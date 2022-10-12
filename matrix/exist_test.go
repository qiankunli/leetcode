package matrix

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(Exist(board, "SEE"))
}
