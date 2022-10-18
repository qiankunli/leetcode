package digit

import (
	"fmt"
	"math"
)

func Reverse(x int) int {
	return reverse(x)
}
func reverse(x int) int {
	negative := 1
	if x < 0 {
		negative = -1
		x = -x
	}
	len := Len(x)
	var sum int64
	r := x
	for i := 0; i < len; i++ {
		c := r % 10
		sum = sum + int64(c*int(math.Pow10(len-i-1)))
		r = r / 10
	}
	if sum > math.MaxInt32 {
		return 0
	}
	return int(sum) * negative
}
func Len(x int) int {
	s := fmt.Sprintf("%d", x)
	return len(s)
}
