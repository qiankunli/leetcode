package string

import "strings"

func LongestCommonChild(s1, s2 string) int {
	return longestCommonChild2(s1, s2)
}

func longestCommonChild(s1, s2 string) int {
	n1 := len(s1)
	n2 := len(s2)
	if n1 == 0 || n2 == 0 {
		return 0
	}
	dp := make([][]int, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1)
	}

	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			s1i := s1[i-1 : i]
			s2j := s2[j-1 : j]
			if strings.Compare(s1i, s2j) == 0 {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[n1][n2]
}

// 最长公共子串 用回溯法就比较难实现，因为回溯法 只能判断字符串的头，但判断长度要看字符串的尾部
// 从后面遍历就成功了
func longestCommonChild2(s1, s2 string) int {
	n1 := len(s1)
	n2 := len(s2)
	if n1 == 0 || n2 == 0 {
		return 0
	}
	s1c := s1[n1-1 : n1]
	s2c := s2[n2-1 : n2]

	if strings.Compare(s1c, s2c) == 0 {
		return 1 + longestCommonChild2(s1[:n1-1], s2[:n2-1])
	}
	l := longestCommonChild2(s1, s2[:n2-1])
	r := longestCommonChild2(s1[:n1-1], s2)
	return max(l, r)

}
