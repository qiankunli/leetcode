package string

import "strings"

// 139
// 第一版回溯法 超时了
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return false
	}
	if len(wordDict) == 0 {
		return false
	}
	return dfs(s, wordDict)
}
func dfs(s string, wordDict []string) bool {
	// 满足 s 都拼完了
	if len(s) == 0 {
		return true
	}
	for _, word := range wordDict {
		if startWith(s, word) { // 此处优先匹配最长的可以优化下，可以先对wordDict 按长度排序一下
			match := dfs(s[len(word):], wordDict)
			if !match {
				continue
			}
			return true
		}
	}
	// 没有单词匹配到
	return false
}

func startWith(s, target string) bool {

	n := len(target)
	if len(s) < n {
		return false
	}
	for i := 0; i < n; i++ {
		sc := s[i : i+1]
		tc := target[i : i+1]
		if strings.Compare(sc, tc) != 0 {
			return false
		}
	}
	return true
}
