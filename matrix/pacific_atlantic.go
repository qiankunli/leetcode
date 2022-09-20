package matrix

// 417
// 太平洋大西洋水流问题
// 要求给出所有解，所以考虑回溯法遍历所有可能路径
// 遍历树要考虑两点：树的根，树有几个分叉，如何记录遍历的结果
func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	if m == 0 {
		return nil
	}
	n := len(heights[0])
	p := make([][]int, m) // 太平洋   0 表示无法到  1 表示可以
	a := make([][]int, m) // 太西洋
	for i := 0; i < m; i++ {
		p[i] = make([]int, n)
		a[i] = make([]int, n)
	}
	dirs = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < m; i++ {
		dfs(heights, p, m, n, i, 0)
	}
	for i := 0; i < m; i++ {
		dfs(heights, a, m, n, i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(heights, p, m, n, 0, j)
	}
	for j := 0; j < n; j++ {
		dfs(heights, a, m, n, m-1, j)
	}
	result := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] == 1 && a[i][j] == 1 {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

var dirs [][]int

func dfs(heights, ocean [][]int, m, n, i, j int) {
	if ocean[i][j] == 1 {
		return
	}
	ocean[i][j] = 1
	for _, dir := range dirs {
		nx := i + dir[0]
		ny := j + dir[1]
		if nx >= 0 && nx < m && ny >= 0 && ny < n && heights[i][j] <= heights[nx][ny] {
			dfs(heights, ocean, m, n, nx, ny)
		}
	}
}
