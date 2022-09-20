package matrix

// 200
func numIslands(grid [][]byte) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(grid [][]byte, m, n, i, j int)
	dfs = func(grid [][]byte, m, n, i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 0 { // 碰到了0 也就是碰到了岛屿的边界
			return
		}
		grid[i][j] = 0
		for _, dir := range dirs {
			nx := i + dir[0]
			ny := j + dir[1]
			dfs(grid, m, n, nx, ny)
		}
	}
	// 从 =1 开始遍历，所有相连的1 置0
	c := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				c++
				dfs(grid, m, n, i, j)
			}
		}
	}
	return c
}
