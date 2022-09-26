package string

import (
	"fmt"
	"strconv"
)

func AddStrings(num1 string, num2 string) string {
	return addStrings(num1, num2)
}

// 415
func addStrings(num1 string, num2 string) string {
	n1 := len(num1)
	n2 := len(num2)

	// 临界条件

	p1 := n1 - 1
	p2 := n2 - 1
	s := ""
	extra := 0
	for p1 >= 0 && p2 >= 0 {
		v, e := sumWithExtra(num1[p1:p1+1], num2[p2:p2+1], extra)
		s = v + s
		extra = e
		p1--
		p2--
	}
	for p1 >= 0 {
		v, e := sumWithExtra(num1[p1:p1+1], "", extra)
		s = v + s
		extra = e
		p1--
	}
	for p2 >= 0 {
		v, e := sumWithExtra(num2[p2:p2+1], "", extra)
		s = v + s
		extra = e
		p2--
	}
	if extra == 1 {
		s = fmt.Sprintf("%d", extra) + s
	}
	return s
}
func sumWithExtra(s1, s2 string, extra int) (string, int) {
	var v1, v2 int
	if len(s1) > 0 {
		v1, _ = strconv.Atoi(s1)
	}
	if len(s2) > 0 {
		v2, _ = strconv.Atoi(s2)
	}
	v := v1 + v2 + extra
	if v >= 10 {
		v = v - 10
		extra = 1
	} else {
		extra = 0
	}
	return fmt.Sprintf("%d", v), extra
}
