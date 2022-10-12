package matrix

// 221
func maximalSquare(matrix [][]byte) int {
	// 二分法，n  先找是否存在 n/2 的正方形，存在则扩大，不存在则缩小
	// 没有想到动态规划

	m := len(matrix)
	n := len(matrix[0])

	dp := make([][]int, m) // 以ij 为右下角 的 最大正方形的边长
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// 初始化
	// 第一个行和第一列都固定为1

	// 状态转移
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min3(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
				}
			}
		}
	}
	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if max < dp[i][j] {
				max = dp[i][j]
			}
		}
	}
	return max * max
}
