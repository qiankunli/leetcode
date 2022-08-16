package array

import (
	"fmt"
	"testing"
)

func TestMaxProfit2(t *testing.T) {
	//prices := []int{7, 1, 5, 3, 6, 4}
	//prices := []int{1, 2, 3, 4, 5}
	prices := []int{3, 2, 6, 5, 0, 3}
	fmt.Println(MaxProfit2(prices))
}
