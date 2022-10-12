package matrix

import (
	"fmt"
	"strings"
)

func Exist(board [][]byte, word string) bool {
	return exist(board, word)
}

func exist2(board [][]byte, word string) bool {
	// 从一个 单元格开始，上下左右搜索，搜索到就继续， 搜索不到就终止

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])

	visited := make([][]bool, m) // 标记访问过的点
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	var dfs func(board [][]byte, word string, k, i, j int) bool
	dfs = func(board [][]byte, word string, k, i, j int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k+1 >= len(word) {
			return true
		}
		visited[i][j] = true
		for _, dir := range dirs {
			x := i + dir[0]
			y := j + dir[1]
			if x < m && y < n && x >= 0 && y >= 0 {
				if visited[x][y] {
					continue
				}
				if dfs(board, word, k+1, x, y) {
					return true
				}
			}
		}
		visited[i][j] = false
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(board, word, 0, i, j) {
					return true
				}
			}
		}
	}
	return false
}

// 超时了
func exist(board [][]byte, word string) bool {
	// 从一个 单元格开始，上下左右搜索，搜索到就继续， 搜索不到就终止

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	path := make([][]int, 0) // 记住走过的路，防止从来路回去，实际可以用 visited[i][j] 替代
	// i j 表示从哪里开始搜索
	var dfs func(board [][]byte, word string, i, j int) bool
	dfs = func(board [][]byte, word string, i, j int) bool {
		l := len(word)
		if l == 0 {
			return true
		}
		c := word[0:1]
		bc := string(board[i][j])
		if strings.Compare(bc, c) != 0 {
			return false
		}
		if l == 1 { // 只有一个字符
			return true
		}
		for _, dir := range dirs {
			x := i + dir[0]
			y := j + dir[1]
			if x < m && y < n && x >= 0 && y >= 0 {
				fmt.Println(word[1:], x, y, m, n)
				if containsEle(path, []int{x, y}) { // 防止从来路回去
					continue
				}
				if dfs(board, word[1:], x, y) {
					return true
				}
			}
		}
		return false
	}

	c := word[0:1]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			bc := string(board[i][j]) // 将数字转为字符串
			fmt.Println("==>", bc, i, j)
			if strings.Compare(bc, c) == 0 {
				if dfs(board, word, i, j) {
					return true
				}
			}
		}
	}
	return false
}

func containsEle(path [][]int, ele []int) bool {
	for _, e := range path {
		if e[0] == ele[0] && e[1] == ele[1] {
			return true
		}
	}
	return false
}
