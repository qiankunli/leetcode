package array

import "math"

func Trap(height []int) int {
	return trap(height)
}

// 42
func trap(height []int) int {
	// 1. 先求出 所有的高峰，所谓高峰，就是左右两侧都比它小
	// 2. 计算两个高峰之间可以存多少水
	// 3. bad case 进一步优化下peaks，如果peaks 发现左右两边比自己更高，则销毁； 或者添加peak的时候，如果发现peaks 最后一个元素 比前一个值小，则覆盖最后一个元素
	n := len(height)

	// 求高峰
	peaks := make([]int, 0)
	for i := 0; i < n; i++ {
		if getHeight(height, i) > getHeight(height, i-1) && getHeight(height, i) > getHeight(height, i+1) {
			peaks = append(peaks, i)
		}
	}
	sum := 0
	for i := 0; i < len(peaks)-1; i++ {
		sum += total(height, peaks[i], peaks[i+1])
	}
	return sum
}
func getHeight(height []int, index int) int {
	if index < 0 || index >= len(height) {
		return math.MinInt
	}
	return height[index]
}
func total(height []int, start, end int) int {
	// 最低点
	min := min(height[start], height[end])
	sum := 0
	for i := start + 1; i < end; i++ {
		sum += min - height[i]
	}
	return sum
}
