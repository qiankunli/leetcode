package array

// 69
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x <= 3 {
		return 1
	}
	if x == 4 {
		return 2
	}
	mid := x / 2
	// x 越大，结果离 i 越近，所以i 从小开始遍历
	for i := 2; i <= mid+1; i++ {
		if i*i > x {
			return i - 1
		}
	}
	return -1
}
