package string

import (
	"fmt"
	"strings"
)

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}

// 5
func longestPalindrome(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return s
	}
	if n == 2 {
		if strings.Compare(s[:1], s[1:]) == 0 {
			return s
		}
		return s[:1]
	}
	// 如果只是求长度，dp 的值可以是长度，但因为要返回真实的字符串，所以还需返回最长位置
	dp := make([][]bool, n)
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, n)
	}

	maxLen := 1 // 最小1个长度还是有的
	begin := 0

	// dp[i][j] = dp[i+1][j-1] && s[i:i+1] == s[j-1:j]
	for l := 2; l <= n; l++ {
		for i := 0; i < n; i++ {
			j := i + l - 1
			if j >= n {
				break
			}
			si := s[i : i+1]
			sj := s[j : j+1]
			// 计算 dp[i][j]
			if strings.Compare(si, sj) != 0 {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

// 第一反应是滑动窗口
func longestPalindrome2(s string) string {
	n := len(s)
	if n <= 1 {
		return ""
	}
	l := 0
	r := 1
	max := ""
	for r < n && l <= r {
		ok := ok(s, l, r)
		fmt.Println(s[l:r+1], ok)
		if !ok {
			r++
		} else {
			if len(max) < r-l+1 {
				max = s[l : r+1]
			}
			r++
		}
	}
	return max
}
func ok(s string, l, r int) bool {
	li := l
	ri := r
	for li < ri {
		ls := s[li : li+1]
		rs := s[ri : ri+1]
		if strings.Compare(ls, rs) != 0 {
			return false
		}
		li++
		ri--
	}
	return true
}
