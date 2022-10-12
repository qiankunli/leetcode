package array

import "math"

// 84
// 超时了，披着dp 外衣的暴力法
func largestRectangleArea(heights []int) int {
	// 滑动窗口，从左到右解决不了
	// dp[i] 以 heights[i] 结尾能划定的矩阵的最大面积，最后实际上暴力法，没省什么
	//

	n := len(heights)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = heights[0]

	for i := 1; i < n; i++ {
		// dp[i] = dp[i-1] 就是不考虑 heights[i] 否则就是考虑 heights[i]
		dp[i] = max(dp[i-1], areaWithTail(heights, i))
	}

	return dp[n-1]

}

// [0,tail] 矩阵的最大面积，且一定要以tail 结尾
func areaWithTail(heights []int, tail int) int {
	m := 0
	for i := tail; i >= 0; i++ {
		a := area(heights, i, tail)
		if m < a {
			m = a
		}
	}
	return m
}
func area(heights []int, head, tail int) int {

	width := tail - head + 1
	height := math.MaxInt // 高度是最小元素
	for i := head; i <= tail; i++ {
		if height > heights[i] {
			height = heights[i]
		}
	}
	return width * height
}
