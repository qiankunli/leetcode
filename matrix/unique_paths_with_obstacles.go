package matrix

func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	return uniquePathsWithObstacles(obstacleGrid)
}

// 63
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 { // 如果起点都是障碍，那就为0
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始化第一行和第一列
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0]
		if obstacleGrid[i][0] == 1 {
			dp[i][0] = 0
		}
	}
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1]
		if obstacleGrid[0][j] == 1 {
			dp[0][j] = 0
		}
	}

	// 初始化其它行列
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			}
		}
	}
	return dp[m-1][n-1]
}

func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 { // 如果起点都是障碍，那就为0
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1 // 这段代码很重要的问题是，对 第一行第一列没有很优雅的办法 根据 obstacleGrid[i][j] 调整，所以最后 只好将初始化化 和递推过程拆开
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			}
		}
	}
	return dp[m-1][n-1]
}

func uniquePathsWithObstacles3(obstacleGrid [][]int) int {
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	if obstacleGrid[0][0] == 1 { // 如果起点都是障碍，那就为0
		return 0
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			}
		}
	}
	return dp[m-1][n-1]
}
