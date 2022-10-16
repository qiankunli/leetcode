package digit

import (
	"fmt"
	"strings"
)

func GetPermutation(n int, k int) string {
	return getPermutation(n, k)
}

// 60
// 想到了dfs 暴力却没想到剪枝
func getPermutation(n int, k int) string {

	var path string
	var dfs func(idx, k int)
	dfs = func(idx, k int) {
		if idx >= n {
			return
		}
		c := f(n - idx - 1) // 在idx +1 层，每个分叉有c 个叶子节点
		for i := 1; i <= n; i++ {
			if strings.Contains(path, fmt.Sprintf("%d", i)) { // 滤掉已经记录过的数字
				continue
			}

			if k <= c {
				path += fmt.Sprintf("%d", i)
				dfs(idx+1, k) // 到下一层考虑下
				break
			}
			k -= c // 如果k 大于 c 就放弃值为i 的分叉
		}
	}
	dfs(0, k)
	return path
}

func f(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return f(n-1) * n
}
