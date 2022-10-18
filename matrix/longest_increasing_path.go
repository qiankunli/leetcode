package matrix

// 329
func longestIncreasingPath(matrix [][]int) int {
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	// 如果使用朴素深度优先搜索，时间复杂度是指数级，会超出时间限制，因此必须加以优化
	// 朴素深度优先搜索的时间复杂度过高的原因是进行了大量的重复计算，同一个单元格会被访问多次，每次访问都要重新计算。由于同一个单元格对应的最长递增路径的长度是固定不变的，因此可以使用记忆化的方法进行优化。
	// cache[i][j] 以ij 为起点的最长递增路径
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
	}

	var dfs func(matrix [][]int, m, n, i, j int) int
	dfs = func(matrix [][]int, m, n, i, j int) int {
		// 如果大于0 说明已经被计算过
		if cache[i][j] > 0 {
			return cache[i][j]
		}
		max := 1 // 最小是1
		for _, dir := range dirs {
			nx := i + dir[0]
			ny := j + dir[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n { // 越界
				continue
			}
			if matrix[i][j] >= matrix[nx][ny] {
				continue
			}
			c := dfs(matrix, m, n, nx, ny)
			cache[nx][ny] = c
			if max < c+1 {
				max = c + 1
			}
		}
		return max
	}

	max := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			c := dfs(matrix, m, n, i, j)
			if max < c {
				max = c
			}
		}
	}
	return max
}
