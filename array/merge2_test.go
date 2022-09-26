package array

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	Merge(nums1, 3, nums2, len(nums2))
	fmt.Println(nums1)
}
