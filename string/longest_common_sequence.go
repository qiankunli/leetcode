package string

import "strings"

func LongestCommonSequence(s1, s2 string) int {
	return longestCommonChild(s1, s2)
}

func longestCommonSequence(s1, s2 string) int {
	n1 := len(s1)
	n2 := len(s2)
	if n1 == 0 || n2 == 0 {
		return 0
	}
	s1c := s1[0:1]
	s2c := s2[0:1]
	if strings.Compare(s1c, s2c) == 0 {
		return 1 + longestCommonSequence(s1[1:], s2[1:])
	}
	l := longestCommonSequence(s1, s2[1:])
	r := longestCommonSequence(s1[1:], s2)
	return max(l, r)
}
