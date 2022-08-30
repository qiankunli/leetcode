package matrix

func KthSmallest(matrix [][]int, k int) int {
	return kthSmallest(matrix, k)
}

// 378
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	l := matrix[0][0]
	r := matrix[n-1][n-1]
	// 第k 小的数是result，那么现在找到某个数 matrix 有k-1个数比它小，那这个数就是result
	// 且这个数一定在 [l,r] 范围内
	for l < r {
		mid := (l + r) / 2
		c := findLessCount(matrix, mid)
		if c < k {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func findLessCount(matrix [][]int, target int) int {
	n := len(matrix)
	i := 0
	j := n - 1
	c := 0
	for i < n && j >= 0 {
		if matrix[i][j] <= target {
			c += j + 1
			i++
		} else {
			j--
		}
	}
	return c

}
