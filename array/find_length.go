package array

// 718

func findLength(nums1 []int, nums2 []int) int {

	// 就是最长公共子串搞法
	// dp[i][j] = dp[i-1][j-1] + 1; nums1[i] = nums2[j]
	// dp[i][j] = max(dp[i-1][j],dp[i][j-1])

	m := len(nums1)
	n := len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	max := 0
	// dp[i][j] 表示以 num1[i] 和 nums[j] 结尾的数组的 公共子数组长度，所以 nums[0] 也是要考虑的
	// 但是公式又是 正推的 有 dp[i-1][j-1] ，就只能故意扩大一位，以便计算
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// dp[i][j] = max(dp[i][j-1],dp[i-1][j])
				dp[i][j] = 0 // 如果以 num1[i] 和 nums[j] 结尾的数组 不一样，最大公共长度直接就为0了
			}
			// 最后只获取最大值
			if max < dp[i][j] {
				max = dp[i][j]
			}
		}
	}
	return max
}
