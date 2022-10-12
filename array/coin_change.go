package array

import (
	"math"
)

func CoinChange(coins []int, amount int) int {
	return coinChange2(coins, amount)
}

// 322
func coinChange(coins []int, amount int) int {

	// dp[0] = 0  是为了后面计算方便。因为是要加1 的
	// dp[i] = math.MaxInt 是为了标记找不到。找不到的dp 不参与到关系转移公式里

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

func coinChange2(coins []int, amount int) int {

	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt
		for _, c := range coins {
			if i >= c && dp[i-c] != math.MaxInt {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
