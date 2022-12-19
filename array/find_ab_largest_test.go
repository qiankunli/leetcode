package array

import (
	"fmt"
	"testing"
)

func TestFindABLargest(t *testing.T) {
	s := []int{1, 3, 40, 100, 4, 30, 5, 9, 7, 6, 8, 2, 18, 16, 37, 2, 0}
	start, end := page(1, 3)
	FindABLargest(s, start, end)
	fmt.Println(s[start:end])
	fmt.Println(s)

	//s = []int{1, 3, 40, 100, 4, 30, 5, 9, 7, 6, 8, 2, 18, 16, 37, 0}
	//start, end = page(2, 3)
	//FindABLargest(s, start, end)
	//fmt.Println(s[start:end])
	//fmt.Println(s)
	//
	//s = []int{1, 3, 40, 100, 4, 30, 5, 9, 7, 6, 8, 2, 18, 16, 37, 0}
	//start, end = page(3, 3)
	//FindABLargest(s, start, end)
	//fmt.Println(s[start:end])
	//fmt.Println(s)
	//
	//s = []int{1, 3, 40, 100, 4, 30, 5, 9, 7, 6, 8, 2, 18, 16, 37, 0}
	//start, end = page(4, 3)
	//FindABLargest(s, start, end)
	//fmt.Println(s[start:end])
	//fmt.Println(s)
	//
	//s = []int{1, 3, 40, 100, 4, 30, 5, 9, 7, 6, 8, 2, 18, 16, 37, 0}
	//start, end = page(5, 3)
	//FindABLargest(s, start, end)
	//fmt.Println(s[start:end])
	//fmt.Println(s)
	//
	//s = []int{1, 3, 40, 100, 4, 30, 5, 9, 7, 6, 99, 8, 2, 18, 16, 37, 0}
	//start, end = page(6, 3)
	//FindABLargest(s, start, end)
	//fmt.Println(s[start:])
	//fmt.Println(s)
}

func page(pageNum, pageSize int) (int, int) {
	start := (pageNum - 1) * pageSize
	return start, start + pageSize
}
