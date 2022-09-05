package array

import "strings"

func GenerateParenthesis(n int) []string {
	return generateParenthesis(n)
}

/**
1. 构造 ((())) 遍历所有可能子序列
2. 二叉树，左节点表示当前选( 右节点表示当前选)
*/
// 22
func generateParenthesis(n int) []string {
	path := ""
	result := make([]string, 0)
	doGenerateParenthesis(n, path, &result)
	return result
}
func doGenerateParenthesis(n int, path string, result *[]string) {

	if len(path) == n*2 {
		*result = append(*result, path)
		return
	}
	cl := count(path, "(")
	cr := count(path, ")")
	if cl < n {
		path += "(" // 当前选(
		doGenerateParenthesis(n, path, result)
		path = path[:len(path)-1]
	}
	// 左括号的数量必须大于右括号的数量才能加右括号
	if cr < n && cl > cr {
		path += ")" // 当前选)
		doGenerateParenthesis(n, path, result)
	}
}

func count(s, target string) int {
	return strings.Count(s, target)
}
