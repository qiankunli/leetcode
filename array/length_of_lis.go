package array

import (
	"math"
)

func LengthOfLIS(nums []int) int {
	return lengthOfLIS3(nums)
}

// 300
// 因为不连续，所以不能用滑动窗口
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		// 因为可以不连续，所以计算dp[i] 要考虑dp[0] 到dp[i-1] 的所有情况
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return dp[n-1]
}

// 没事手搓一下
func lengthOfLIS3(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {

		// 找到 [0,i) 的dp 最大值
		max := 0
		for j := 0; j < i; j++ {
			if max < dp[j] && nums[i] > nums[j] {
				max = dp[j]
			}
		}
		dp[i] = max + 1
	}
	// dp[i] 表示以nums[i] 结尾的最长子序列的最大值，真正的最大值要遍历看下
	max := 0
	for i := 0; i < n; i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	return max
}

// 超时了
// 回溯法本质就是暴力遍历
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	path2 = make([]int, 0)
	return doLengthOfLIS(nums, 0)
}

var path2 []int

func doLengthOfLIS(nums []int, start int) int {
	if start >= len(nums) {
		return len(path2)
	}
	var last int
	if len(path2) == 0 {
		last = math.MinInt
	} else {
		last = path2[len(path2)-1]
	}
	if nums[start] > last {
		// 使用当前值
		path2 = append(path2, nums[start])
		l1 := doLengthOfLIS(nums, start+1)
		path2 = path2[:len(path2)-1]
		// 不用当前值
		l2 := doLengthOfLIS(nums, start+1)
		if l1 > l2 {
			return l1
		}
		return l2
	} else {
		return doLengthOfLIS(nums, start+1)
	}
}
