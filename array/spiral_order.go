package array

import "fmt"

func SpiralOrder(matrix [][]int) []int {
	return spiralOrder(matrix)
}

var path []int

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return nil
	}
	columns := len(matrix[0])
	path = make([]int, 0)
	doSpiralOrder(matrix, 0, 0, rows, columns)
	return path
}
func doSpiralOrder(matrix [][]int, row, column int, rows, columns int) {
	if rows == 0 || columns == 0 {
		return
	}
	for i := column; i < column+columns; i++ {
		path = append(path, matrix[row][i])
	}
	fmt.Println(path)
	for i := row + 1; i < row+rows; i++ {
		path = append(path, matrix[i][column+columns-1])
	}
	fmt.Println(path)
	// 大于1行1列，可以回旋一下
	if rows > 1 && columns > 1 {
		for i := column + columns - 2; i >= column; i-- {
			path = append(path, matrix[row+rows-1][i])
		}
		fmt.Println(path)
		for i := row + rows - 2; i > row; i-- {
			path = append(path, matrix[i][column])
		}
		fmt.Println(path)
	}

	if rows > 2 && columns > 2 {
		doSpiralOrder(matrix, row+1, column+1, rows-2, columns-2)
	}
}

func spiralOrder2(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return nil
	}
	columns := len(matrix[0])
	color := make([][]int, rows)
	for i := 0; i < rows; i++ {
		color[i] = make([]int, columns)
	}
	i := 0
	j := 0
	c := 1
	n := rows * columns
	result := make([]int, 0)
	result = append(result, matrix[0][0])
	color[0][0] = 1
	for c < n {
		// 向右
		if j+1 < columns && color[i][j+1] == 0 {
			result = append(result, matrix[i][j+1])
			color[i][j+1] = 1
			c++
			j++
			continue
		}
		// 向下
		if i+1 < rows && color[i+1][j] == 0 {
			result = append(result, matrix[i+1][j])
			color[i+1][j] = 1
			c++
			i++
			continue
		}
		// 向左
		if j-1 >= 0 && color[i][j-1] == 0 {
			result = append(result, matrix[i][j-1])
			color[i][j-1] = 1
			c++
			j--
			continue
		}
		// 向上
		if i-1 >= 0 && color[i-1][j] == 0 {
			result = append(result, matrix[i-1][j])
			color[i-1][j] = 1
			c++
			i--
			continue
		}
	}
	return result
}
