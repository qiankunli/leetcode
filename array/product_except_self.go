package array

// 238
func productExceptSelf(nums []int) []int {
	// 暴力法是 n^2
	// 这种场景 一般是多次遍历，left[i]  表示 0 到i 的乘积，right[i] 表示 i 到n 的乘积
	// multi[i] = left[i-1] * right[i+1]
	// 虽然是自己想的，但有点脑筋急转弯了

	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)

	left[0] = nums[0]
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i]
	}

	right[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i]
	}

	result := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			result[i] = right[1]
		} else if i == n-1 {
			result[i] = left[n-2]
		} else {
			result[i] = left[i-1] * right[i+1]
		}
	}
	return result
}
