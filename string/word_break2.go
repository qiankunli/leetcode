package string

import "strings"

// 140

func WordBreak2(s string, wordDict []string) []string {
	return wordBreak2(s, wordDict)
}

func wordBreak2(s string, wordDict []string) []string {
	if len(s) == 0 {
		return nil
	}
	if len(wordDict) == 0 {
		return nil
	}
	path := make([]string, 0)
	result := make([][]string, 0)
	var dfs func(s string, wordDict []string, path []string) bool
	dfs = func(s string, wordDict []string, path []string) bool {
		// 满足 s 都拼完了
		if len(s) == 0 {
			pathCopy := make([]string, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return true
		}
		for _, word := range wordDict {
			if strings.HasPrefix(s, word) {
				path = append(path, word)
				match := dfs(s[len(word):], wordDict, path)
				path = path[:len(path)-1] // 因为访问的所有结果，所以不管是否匹配都要回溯，记录是否匹配是为了作为终止条件
				if !match {
					continue
				}
			}
		}
		// 没有单词匹配到
		return false
	}
	dfs(s, wordDict, path)
	sentences := make([]string, 0)
	for _, words := range result {
		sentence := strings.Join(words, " ")
		sentences = append(sentences, sentence)
	}
	return sentences
}
