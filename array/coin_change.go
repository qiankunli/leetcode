package array

import (
	"math"
)

func CoinChange(coins []int, amount int) int {
	return coinChange(coins, amount)
}

// 322
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for _, c := range coins {
			// dp[i-c] = math.MaxInt 表示还 没有方案凑到 amount = i-c
			if c <= i && dp[i-c] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
