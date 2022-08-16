package array

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	fmt.Println(CoinChange([]int{1, 2, 5}, 11))
	fmt.Println(CoinChange([]int{2}, 3))
	fmt.Println(CoinChange([]int{1}, 0))
}
