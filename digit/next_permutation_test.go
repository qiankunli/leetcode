package digit

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	nums := []int{1, 2, 4, 3}
	NextPermutation(nums)
	fmt.Println(nums)
}
