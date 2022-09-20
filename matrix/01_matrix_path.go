package matrix

func MatrixPath(matrix [][]int) [][]int {
	return matrixPath(matrix)
}

var c0 int
var m, n int
var path, result, visited [][]int

func matrixPath(matrix [][]int) [][]int {
	m = len(matrix)
	if m == 0 {
		return nil
	}
	n = len(matrix[0])
	// 统计0的个数
	c0 = count(matrix, 0)
	path = make([][]int, 0)
	visited = make([][]int, m)
	for i := 0; i < n; i++ {
		visited[i] = make([]int, n)
	}
	result = make([][]int, c0)
	path = append(path, []int{0, 0})
	visited[0][0] = 1
	doMatrixPath(matrix, 0, 0, path)
	return result
}

// path[i][j] = 1 记录不了路径
func doMatrixPath(matrix [][]int, i, j int, path [][]int) {
	if i == 2 && j == 2 {

	}
	if i < 0 || i >= m {
		return
	}
	if j < 0 || j >= n {
		return
	}
	if i == m-1 && j == n-1 {
		if c0 == len(path) {
			copy(result, path)
		}
		return
	}

	if matrix[i][j] == 1 {
		return
	}
	// matrix[i][j] == 0
	// 向下
	if i < m-1 && visited[i+1][j] != 1 { // 没被访问过
		path = append(path, []int{i + 1, j})
		visited[i+1][j] = 1
		doMatrixPath(matrix, i+1, j, path)
		visited[i+1][j] = 0
		path = path[:len(path)-1]
	}
	// 向右
	if j < n-1 && visited[i][j+1] != 1 { // 没被访问过
		path = append(path, []int{i, j + 1})
		visited[i][j+1] = 1
		doMatrixPath(matrix, i, j+1, path)
		visited[i][j+1] = 0
		path = path[:len(path)-1]
	}
	// 向上
	if i > 0 && visited[i-1][j] != 1 { // 没被访问过
		path = append(path, []int{i - 1, j})
		visited[i-1][j] = 1
		doMatrixPath(matrix, i-1, j, path)
		visited[i-1][j] = 0
		path = path[:len(path)-1]
	}
	// 向左
	if j > 0 && visited[i][j-1] != 1 { // 没被访问过
		path = append(path, []int{i, j - 1})
		visited[i][j-1] = 1
		doMatrixPath(matrix, i, j-1, path)
		visited[i][j-1] = 0
		path = path[:len(path)-1]
	}

}
func count(matrix [][]int, target int) int {
	c := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == target {
				c++
			}
		}
	}
	return c
}
