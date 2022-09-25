package string

import (
	"fmt"
	"strings"
)

func LengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring3(s)
}

// 3

func lengthOfLongestSubstring3(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	l := 0
	r := 0
	max := 0
	cache := make(map[string]int, 0)

	for l <= r && r < n {
		cr := s[r : r+1]
		for cache[cr] > 0 { // l移到重复的字符后面
			cl := s[l : l+1]
			cache[cl]--
			l++
		}
		cache[cr] = 1
		if max < r-l+1 {
			max = r - l + 1
		}
		r++
	}

	return max

}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	if len(s) == 2 {
		if strings.Compare(s[:1], s[1:]) == 0 {
			return 1
		} else {
			return 2
		}
	}
	l := 0
	r := 1
	result := 1
	for r < len(s) && l <= r {
		if have(s, l, r) {
			if result < r-l+1 {
				result = r - l + 1
			}
			r++
		} else {
			l++
		}
	}
	return result
}
func have(s string, l int, r int) bool {
	hash := make(map[string]bool)
	for j := l; j <= r; j++ {
		s1 := s[j : j+1]
		if _, ok := hash[s1]; ok {
			return false
		}
		hash[s1] = true
	}
	return true
}

// 3
// 一开始没有反应过来这是一个滑动窗口的题目
func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	if len(s) == 2 {
		if strings.Compare(s[:1], s[1:]) == 0 {
			return 1
		} else {
			return 2
		}
	}
	l := 0
	r := len(s) - 1
	result := 0
	for l <= r {
		mid := (l + r) / 2
		have := findLen(s, mid+1)
		fmt.Println(l, r, mid, s, have)
		if have == -1 {
			r = mid - 1
		} else {
			result = mid + 1
			l = mid + 1
		}
	}
	return result
}
func findLen(s string, n int) int {
	l := len(s)
	for i := 0; i <= l-n; i++ {
		hash := make(map[string]bool)
		fmt.Println(s[i : i+n])
		j := i
		for j < i+n {
			s1 := s[j : j+1]
			fmt.Println(s1)
			if _, ok := hash[s1]; ok {
				break
			}
			hash[s1] = true
			j++
		}
		if j == i+n {
			return 0
		}
	}
	return -1
}
