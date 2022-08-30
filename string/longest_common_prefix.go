package string

func LongestCommonPrefix(strs []string) string {
	return longestCommonPrefix(strs)
}

// 14
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	common := ""
	firstStr := strs[0]
	curIndex := 0
	for {
		if curIndex >= len(firstStr) {
			return common
		}
		cur := firstStr[curIndex]
		for i := 1; i < len(strs); i++ {
			s := strs[i]
			if curIndex >= len(s) {
				return common
			}
			if cur != s[curIndex] {
				return common
			}
		}
		common += string(cur)
		curIndex++
	}
	return common
}
