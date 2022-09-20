package array

import (
	"fmt"
	"testing"
)

func TestLengthOfLIS(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(LengthOfLIS(nums))

	//nums := []int{0, 1, 0, 3, 2, 3}
	//fmt.Println(LengthOfLIS(nums))

	//nums := []int{7, 7, 7, 7, 7, 7, 7}
	//fmt.Println(LengthOfLIS(nums))

}
