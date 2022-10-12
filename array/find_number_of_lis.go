package array

import "fmt"

// 673
func findNumberOfLIS(nums []int) int {
	// 暴力所有子序列
	// 先求出来最长递增子序列的dp，求出最大值，求出个数，不行
	// 知道最大长度的情况下，如果长度是1 ，则直接返回元素的个数，否则求 重复元素个数

	n := len(nums)
	dp := make([]int, n)    // 以 nums[i] 结尾的最长子序列的长度
	count := make([]int, n) // 以 nums[i] 结尾的最长子序列的路径 个数
	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] { // 第一次碰到
					count[i] = count[j]
					dp[i] = dp[j] + 1
				} else if dp[j]+1 == dp[i] { // 第二次或更多次碰到
					count[i] += count[j]
				}
			}
		}
	}
	fmt.Println(dp)
	m := 0 // 找出dp 里最大的值
	for i := 0; i < n; i++ {
		if m < dp[i] {
			m = dp[i]
		}
	}
	c := 0 // 找出最大的值 求和
	for i := 0; i < n; i++ {
		if dp[i] == m {
			c += count[i]
		}
	}
	return c
}
