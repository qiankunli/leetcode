package array

// 11
func maxArea(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	l := 0
	r := n - 1
	max := 0
	for l < r {
		min := minHeight(height, l, r)
		area := height[min] * (r - l)
		if max < area {
			max = area
		}
		if l == min {
			l++
		} else {
			r--
		}
	}
	return max
}
func minHeight(height []int, l, r int) int {
	if height[l] < height[r] {
		return l
	}
	return r
}
