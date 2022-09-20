package array

// 198
func Rob(nums []int) int {
	path := make([]int, 0)
	return rob2(nums, 0, &path)
}
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, n+1)
	// 只能初始化到n-1，因为dp[n-1] 只能是 nums[n-1]，但是 dp[n-2] 可以是 nums[n-2] 也可以是nums[n-1]
	dp[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		dp[i] = max(dp[i+1], dp[i+2]+nums[i])
	}
	return dp[0]
}

func rob2(nums []int, idx int, path *[]int) int {
	n := len(nums)
	if idx >= n {
		return sum(*path)
	}
	*path = append(*path, nums[idx])
	amount1 := rob2(nums, idx+2, path) // 偷当前的
	p := *path
	*path = p[:len(p)-1]
	amount2 := rob2(nums, idx+1, path) // 不偷当前的
	return max(amount1, amount2)
}
