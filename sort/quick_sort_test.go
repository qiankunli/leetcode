package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	values := []int{5, 1, 8, 4, 3, 9}
	QuickSort(values)
	fmt.Println(values)
}
