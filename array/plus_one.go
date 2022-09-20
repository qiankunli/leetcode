package array

// 66
func plusOne(digits []int) []int {
	array := make([]int, 0)
	// 先反转，因为slice 向右边好扩展
	for i := len(digits) - 1; i >= 0; i-- {
		array = append(array, digits[i])
	}
	next := 0 // 记录进位
	for i := 0; i < len(array); i++ {
		sum := array[i] + 1
		if sum < 10 {
			array[i] = sum
			next = 0
			break
		} else {
			array[i] = 0
			next = 1
		}
	}
	if next == 1 {
		array = append(array, next)
	}
	// 再反转回来
	result := make([]int, 0)
	for i := len(array) - 1; i >= 0; i-- {
		result = append(result, array[i])
	}
	return result
}
