package array

import "sort"

func CanPartition(nums []int) bool {
	return canPartition3(nums)
}

// 416
// 求和，拿到平均值，然后找到 是否有部分解可以拼平均值。回溯超时，动态规划优化
// 再进一步 本质是一个在数组中，求和的问题，窗口

func canPartition3(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	sum := sum(nums)
	if sum%2 > 0 {
		return false
	}
	target := sum / 2
	sort.Ints(nums)
	l := 0
	r := 0
	total := 0
	for r < n && l <= n {
		if total == target {
			return true
		}
		if total < target {
			r++
			total += nums[r]
		} else {
			total -= nums[l]
			l--
		}
	}
	return false
}

// 这个方法解决不了   dp[i][j] 只能统计连续和
func canPartition2(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	sum := sum(nums)
	if sum%2 > 0 {
		return false
	}
	dp := make([][]int, n)
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, n)
	}
	for l := 1; l < n; l++ {
		for i := 0; i < n; i++ {
			j := i + l - 1
			if j >= n {
				continue
			}
			if l == 1 {
				dp[i][j] = nums[i]

			} else {
				dp[i][j] = dp[i][j-1] + nums[j]
			}
			if dp[i][j] == sum/2 {
				return true
			}
		}
	}
	return false
}

// 超时
func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}
	sum := sum(nums)
	if sum%2 > 0 {
		return false
	}
	var dfs func(nums []int, idx, target int) bool
	dfs = func(nums []int, idx, target int) bool {
		if idx >= len(nums) {
			return false
		}
		if target == nums[idx] {
			return true
		}
		// 选用当前节点
		b := dfs(nums, idx+1, target-nums[idx])
		if b {
			return true
		}
		return dfs(nums, idx+1, target)
	}
	return dfs(nums, 0, sum/2)
}
