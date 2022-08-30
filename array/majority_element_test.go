package array

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	nums := []int{3, 2, 3}
	fmt.Println(MajorityElement(nums))

	//nums := []int{2, 2, 1, 1, 1, 2, 2}
	//fmt.Println(MajorityElement(nums))
}
