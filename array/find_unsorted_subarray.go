package array

// 581
func findUnsortedSubarray(nums []int) int {
	// 找到一个区域，左边升序，右边升序  中间是乱序       xx 最小值    乱序     最大值 升序
	// 从两侧逼近， 问题是 [2,6,4,8,10,9,15]  左边26   右边915    范围不对

	// 从左向右  left[i]  表示 0 到i 最大值
	// 从右向左  right[i]  表示 i 到n-1 最小值

	// 从左到右  if nums[i] < right[i+1]  继续，否则停止。找到了区域的 左边界

	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)

	// 构建left
	left[0] = nums[0]
	for i := 1; i < n; i++ {
		left[i] = max(left[i-1], nums[i])
	}

	// 构建right
	right[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = min(right[i+1], nums[i])
	}

	// 寻找左边界
	l := 0
	for l < n-1 && nums[l] <= min(right[l+1], nums[l+1]) {
		l++
	}
	// 寻找右边界
	r := n - 1
	for r > 0 && nums[r] >= max(left[r-1], nums[r-1]) {
		r--
	}
	if r <= l {
		return 0
	}
	return r - l + 1
}
