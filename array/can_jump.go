package array

// 55
func canJump(nums []int) bool {

	// dp[0] = true，之后从dp [0] 正推
	n := len(nums)
	if n == 0 {
		return false
	}
	dp := make([]bool, n)
	dp[0] = true
	for i := 0; i < n; i++ {
		if !dp[i] {
			continue
		}
		for j := 0; j <= nums[i]; j++ {
			if i+j >= n {
				continue
			}
			dp[i+j] = true
		}
	}
	return dp[n-1]
}
