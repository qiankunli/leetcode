package string

import "strings"

func ValidPalindrome(s string) bool {
	return validPalindrome(s)
}

// 680
func validPalindrome(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	if n == 1 {
		return true
	}
	return doValidPalindrome(s, 0, n-1)
}
func doValidPalindrome(s string, l, r int) bool {
	if l > r {
		return false
	}
	for l < r {
		sl := s[l : l+1]
		sr := s[r : r+1]
		if strings.Compare(sl, sr) != 0 {
			return doValidPalindrome(s, l+1, r) || doValidPalindrome(s, l, r-1)
		}
		l++
		r--
	}
	return true
}
