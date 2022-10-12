package matrix

// 74
// 第二种方法是两次二分，针对第一列，先找到最接近 target的数，然后在所在行二分查找
func searchMatrix(matrix [][]int, target int) bool {
	// 二分，怎么表示mid
	// 还按一维 数组那么做，mid 还按原来的表示， 映射到二维 的坐标

	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	len := m * n

	l := 0
	r := len - 1
	for l <= r {
		mid := (l + r) / 2
		x := mid / n
		y := mid % n
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
