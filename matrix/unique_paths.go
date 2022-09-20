package matrix

// 62
// 动态规划
func uniquePaths(m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// dp[i][j]   ij 到mn 的路径数
	// dp[i][j]= dp[i+1][j] + dp[i][j+1]
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 {
				dp[i][j] = 1 // 边缘的只有一条路
				continue
			}
			if j == n-1 {
				dp[i][j] = 1 // 边缘的只有一条路
				continue
			}
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}
	return dp[0][0]
}

// 回溯法
// 做选择
func uniquePaths2(m int, n int) int {
	return doUniquePaths(0, 0, m, n)
}
func doUniquePaths(i, j, m int, n int) int {
	if i == m-1 && j == n-1 {
		return 0
	}
	if i == m-2 && j == n-1 {
		return 1
	}
	if i == m-1 && j == n-2 {
		return 1
	}
	down, right := 0, 0
	// 向下
	if i < m-1 {
		down = doUniquePaths(i+1, j, m, n)
	}
	// 向右
	if j < n-1 {
		right = doUniquePaths(i, j+1, m, n)
	}
	return down + right
}
