package string

import (
	"strings"
)

// 最长公共字串
func LongestCommonChildString(s1, s2 string) string {
	return longestCommonChildString(s1, s2)
}

func longestCommonChildString(s1, s2 string) string {
	n1 := len(s1)
	n2 := len(s2)
	if n1 == 0 || n2 == 0 {
		return ""
	}
	// dp[i][j] 表示 s1 前i个字符 s2 前j个字符 的最长公共子串
	// dp[i][j] = dp[i-1][j-1] + 1 ; max(dp[i-1][j], dp[i][j-1])
	// ij 必须从1 开始，但是dp[0][0] 也是必须赋值的，所以长度就加了1
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}
	maxLen := 0
	maxString := ""

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			s1i := s1[i-1 : i]
			s2j := s2[j-1 : j]
			if strings.Compare(s1i, s2j) == 0 {
				dp[i][j] = dp[i-1][j-1] + 1
				if maxLen < dp[i][j] {
					maxLen = dp[i][j]
					maxString = s1[:i]
				}
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return maxString[len(maxString)-maxLen:]
}
