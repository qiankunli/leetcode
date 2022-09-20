package array

import "math"

// 1155

func numRollsToTarget2(n int, k int, target int) int {
	if target == 0 {
		return 0
	}
	limit := int(math.Pow10(9)) + 7 // 取模的值
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, target+1)
	}
	for i := 0; i < n; i++ {
		end := max(target, (i+1)*k) // 大于target 不需要计算，大于 (i+1)*k 骰子拼不出来
		for j := i + 1; j <= end; j++ {
			if i == 0 { // 只有1个骰子 的时候只有一个方法
				dp[i][j] = 1
				continue
			}
			sum := 0
			for c := 1; c <= k; c++ {
				if j < c { // 为了下面的j-c
					continue
				}
				sum += dp[i-1][j-c]
				if sum > limit {
					sum = sum % limit
				}
			}
			dp[i][j] = sum
		}
	}
	return dp[n-1][target]
}

// 超时
func numRollsToTarget(n int, k int, target int) int {
	if target <= 0 { // 这个路径不通
		return 0
	}
	if n == 1 {
		if target > k {
			return 0
		}
		return 1
	}
	sum := 0
	for i := 1; i <= k; i++ {
		sum += numRollsToTarget(n-1, k, target-i)
	}
	return sum
}
