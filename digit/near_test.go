package digit

import (
	"fmt"
	"testing"
)

func TestNear(t *testing.T) {
	nums := []int{9, 3, 4}
	fmt.Println(Near(nums, 23231))
}
