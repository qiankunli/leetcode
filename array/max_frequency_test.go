package array

import (
	"fmt"
	"testing"
)

func TestMaxFrequency(t *testing.T) {
	//nums := []int{1, 2, 4}
	//fmt.Println(MaxFrequency(nums, 5))

	//nums := []int{1, 4, 8, 13}
	//fmt.Println(MaxFrequency(nums, 5))

	//nums := []int{3, 9, 6}
	//fmt.Println(MaxFrequency(nums, 2))

	nums := []int{9940, 9995, 9944, 9937, 9941, 9952, 9907, 9952, 9987, 9964, 9940, 9914, 9941, 9933, 9912, 9934, 9980, 9907, 9980, 9944, 9910, 9997}
	fmt.Println(MaxFrequency(nums, 7925))
}
