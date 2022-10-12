package matrix

func searchMatrix2(matrix [][]int, target int) bool {
	// 之前的题型 可以全局二分， 或者两次二分
	// 看了题解，是从右上角开始，这个思路太难想了
	// 矩阵一个比较直观的二分是，先对比矩阵

	m := len(matrix)
	n := len(matrix[0])

	row := 0
	col := n - 1
	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true
		}
		if target > matrix[row][col] {
			row++
		} else {
			col--
		}

	}
	return false
}
