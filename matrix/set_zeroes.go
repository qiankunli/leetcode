package matrix

import "math"

func SetZeroes(matrix [][]int) {
	setZeroes(matrix)
}

// 73
// 难点就是 怎么区分这里本来就是0 还是后面置0，额外记录行列号
// 空间复杂度 m+n，如果有一个元素为0， 则正好记录为0的行列，之后再将其归零
func setZeroes(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	columns := len(matrix[0])
	// 记录第一个0点的 行号列号，用这个行和列 用来存储其它为0的行号、列号
	r := math.MinInt
	c := math.MinInt
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {

			if matrix[i][j] != 0 {
				continue
			}
			if r == -1 && c == -1 {
				r = i
				c = j
			}
			matrix[i][c] = 0 // 表示i 行为0
			matrix[r][j] = 0 // 表示j 列为0
		}
	}
	if r == math.MinInt && c == math.MinInt { // 说明没有0元素
		return
	}
	// 把 rc 记录的 行、列置为0，rc 先不管
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if r == i || c == j {
				continue
			}
			if matrix[i][c] == 0 || matrix[r][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	// 将 rc 本身置为0
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if r == i || c == j {
				matrix[i][j] = 0
			}
		}
	}
}

// 空间复杂度 m+n
func setZeroes2(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	columns := len(matrix[0])
	blackRows := make(map[int]bool)
	blackColumns := make(map[int]bool)
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if blackRows[i] || blackColumns[j] {
				continue
			}
			if matrix[i][j] > 0 {
				continue
			}
			blackRows[i] = true
			blackColumns[j] = true
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if blackRows[i] || blackColumns[j] {
				matrix[i][j] = 0
			}
		}
	}
}
