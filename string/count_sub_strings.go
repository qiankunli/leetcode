package string

// 647
func countSubstrings(s string) int {
	// 暴力，遍历所有子串，依次判断其是否是回文
	// 之前有最长回文子串的题目，用dp，这个题目本质是统计 dp > 0 的数量。注意的点是 字符可能有重复的，不过题目不做要求

	// dp[i][j]  i 到 j 是否是回文
	// dp[i][j]  = dp[i-1][j+1] + 2 if nums[i] == nums[j]  或者false

	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true // dp 初始化
	}

	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			j := i + l - 1
			if j >= n {
				continue
			}
			if s[i] != s[j] {
				dp[i][j] = false
				continue
			}
			if l == 2 {
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1]
			}
		}
	}

	c := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] {
				c++
			}
		}
	}
	return c
}
