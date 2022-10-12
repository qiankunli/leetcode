package array

// 739
func dailyTemperatures(temperatures []int) []int {
	// 用 right[i] 记录 i 到n-1 的最大值
	// index[i]  记录 right[i] 对应的坐标
	// 尴尬之处在于 这个方法 记录的是最高气温，而不是更高气温
	// 题解答案是单调栈解决

	n := len(temperatures)
	if n == 1 {
		return []int{0}
	}

	right := make([]int, n)
	index := make([]int, n)

	right[n-1] = temperatures[n-1]
	index[n-1] = n - 1
	for i := n - 2; i >= 0; i-- {
		if right[i+1] > temperatures[i] {
			right[i] = right[i+1]
			index[i] = index[i+1]
		} else {
			right[i] = temperatures[i]
			index[i] = i
		}
		// right[i] = max(right[i+1],temperatures[i])
	}
	// fmt.Println(right)
	// fmt.Println(index)

	result := make([]int, n)
	result[n-1] = 0
	for i := 0; i < n-1; i++ {
		if temperatures[i] >= right[i+1] {
			result[i] = 0
		} else {
			result[i] = index[i+1] - i
			for j := i + 1; j < index[i+1]; j++ {
				if temperatures[j] > temperatures[i] {
					result[i] = j - i
					break
				}
			}
		}
	}
	return result
}
