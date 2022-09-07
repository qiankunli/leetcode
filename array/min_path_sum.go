package array

import "fmt"

func MinPathSum(grid [][]int) int {
	return minPathSum(grid)
}

// 64
func minPathSum(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	if rows == 1 {
		return sum(grid[0])
	}
	columns := len(grid[0])
	dp := make([][]int, rows+1)
	for i := 0; i <= rows; i++ {
		dp[i] = make([]int, columns+1)
	}
	for i := 1; i <= rows; i++ {
		for j := 1; j <= columns; j++ {
			if i == 1 {
				dp[i][j] = dp[i][j-1] + grid[i-1][j-1]
				continue
			}
			if j == 1 {
				dp[i][j] = dp[i-1][j] + grid[i-1][j-1]
				continue
			}
			dp[i][j] = min(dp[i][j-1]+grid[i-1][j-1], dp[i-1][j]+grid[i-1][j-1])
		}
	}
	return dp[rows][columns]
}

func minPathSum2(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	if rows == 1 {
		return sum(grid[0])
	}
	path := make([]int, 0)
	path = append(path, grid[0][0])
	return doMinPathSum(grid, &path, 0, 0)
}
func doMinPathSum(grid [][]int, p *[]int, r, c int) int {
	rows := len(grid)
	columns := len(grid[0])
	fmt.Println(*p)
	if r == rows-1 && c == columns-1 {
		return sum(*p)
	}
	// 只能向右
	if r == rows-1 {
		*p = append(*p, grid[r][c+1])
		s := doMinPathSum(grid, p, r, c+1)
		path := *p
		*p = path[:len(*p)-1]
		return s
	}
	// 只能向下
	if c == columns-1 {
		*p = append(*p, grid[r+1][c])
		s := doMinPathSum(grid, p, r+1, c)
		path := *p
		*p = path[:len(*p)-1]
		return s
	}
	*p = append(*p, grid[r][c+1])
	right_sum := doMinPathSum(grid, p, r, c+1)
	path := *p
	*p = path[:len(*p)-1]

	*p = append(*p, grid[r+1][c])
	down_sum := doMinPathSum(grid, p, r+1, c)
	path = *p
	*p = path[:len(*p)-1]

	if right_sum > down_sum {
		return down_sum
	}
	return right_sum
}
