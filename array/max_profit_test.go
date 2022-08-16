package array

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(MaxProfit(prices))
}
