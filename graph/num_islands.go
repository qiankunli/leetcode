package graph

func NumIslands(grid [][]byte) int {
	return numIslands(grid)
}

// 200
func numIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	visited := make([][]byte, rows)
	columns := len(grid[0])
	for i := 0; i < rows; i++ {
		visited[i] = make([]byte, columns)
	}
	num := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if grid[i][j] == 1 && visited[i][j] == 0 {
				if dfs(grid, visited, i, j) > 0 {
					num++
				}
			}
		}
	}
	return num
}
func dfs(grid, visited [][]byte, r, c int) int {
	// 是否遍历过或到了边界
	rows := len(grid)
	columns := len(grid[0])
	// 合理范围
	if r < 0 || c < 0 || r >= rows || c >= columns {
		return 0
	}
	if visited[r][c] == 1 { // 访问过
		return 0
	}
	if grid[r][c] == 0 { // 是否陆地
		return 0
	}
	visited[r][c] = 1
	left := dfs(grid, visited, r, c-1)
	right := dfs(grid, visited, r, c+1)
	up := dfs(grid, visited, r+1, c)
	down := dfs(grid, visited, r-1, c)
	return 1 + left + right + up + down
}
