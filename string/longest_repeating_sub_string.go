package string

import (
	"strings"
)

func LongestRepeatingSubstring(s string) int {
	return longestRepeatingSubstring(s)
}

/*
之前一直想的是去 找最长的，然后算长度，没有反向思考

我们可以将原问题拆分为两个子问题，然后在较优的时间复杂度内依次解决它们。
1. 在 1 到 N 中枚举子串的长度；
2. 对于枚举的长度 L，检查字符串中是否有长度为 L 的重复子串。
*/
// 1062
func longestRepeatingSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}
	l := 0
	r := n - 1
	var result, mid int
	for l <= r {
		mid = (l + r) / 2
		if search(s, mid) == -1 {
			r = mid - 1
		} else {
			result = mid
			l = mid + 1
		}

	}
	return result
}

func search(s string, n int) int {
	l := len(s)
	hash := make(map[string]string, 0)
	for i := 0; i < l-n+1; i++ {
		s1 := s[i : i+n]
		if _, ok := hash[s1]; ok {
			return 0
		}
		hash[s1] = ""
	}
	return -1
}

func longestRepeatingSubstring3(s string) int {
	n := len(s)
	if n <= 1 {
		return 0
	}
	dp := make([][2]int, n+1)
	dp[n-1][1] = 0
	dp[n-1][0] = 0
	dp[n][1] = 0
	dp[n][0] = 0
	// 本质仍是暴力法
	for i := n - 1; i >= 0; i-- {
		dp[i][1] = LongestRepeatingSubstringWithPrefix(s, i)
		dp[i][0] = max(dp[i][1], dp[i+1][0])
	}
	return dp[0][0]
}

func longestRepeatingSubstring2(s string) int {
	if len(s) <= 1 {
		return 0
	}
	return doLongestRepeatingSubstring(s, 0)
}
func doLongestRepeatingSubstring(s string, start int) int {
	if len(s)-start <= 1 {
		return 0
	}
	len1 := doLongestRepeatingSubstring(s, start+1)
	len2 := LongestRepeatingSubstringWithPrefix(s, start)
	return max(len1, len2)
}

// 确定以start 开头的字串，有没有重复字串
func LongestRepeatingSubstringWithPrefix(s string, start int) int {
	// 从start 开始，找到下一个 s[index] = s[start] 然后看看有多长
	max := 0
	for i := start + 1; i < len(s); i++ {
		s1 := s[start : start+1]
		s2 := s[i : i+1]
		if strings.Compare(s1, s2) != 0 {
			continue
		}
		for j := i; j < len(s); j++ {
			s1 := s[start : start+j-i+1]
			s2 := s[i : j+1]
			if strings.Compare(s1, s2) == 0 {
				if max < j-i+1 {
					max = j - i + 1
				}
			}
		}
	}
	return max
}
func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}
