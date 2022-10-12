package array

import (
	"fmt"
	"sort"
)

// 560

// 看了题解，答案是前缀和
func subarraySum3(nums []int, k int) int {
	n := len(nums)
	count := 0
	preSum := make([]int, n+1)
	// preSum[i] 表示  前 i-1 个元素的和
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			// i 到 j 的和
			s := preSum[j+1] - preSum[i]
			if s == k {
				count++
			}
		}

	}

	return count
}

// dp 空间分配超出要求
func subarraySum2(nums []int, k int) int {
	n := len(nums)
	count := 0
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {

		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = nums[j]
			} else {
				dp[i][j] = dp[i][j-1] + nums[j]
			}
			if dp[i][j] == k {
				count++
			}
		}
	}
	return count
}

func subarraySum(nums []int, k int) int {
	// 暴力遍历所有子数组，题目要求是连续的，所以不能排序
	// 重复的地方就在于说，很多地方重复计算了
	n := len(nums)
	sort.Ints(nums)
	fmt.Println(nums)
	count := 0
	for i := 0; i < n; i++ {
		s := 0
		for j := i; j < n; j++ {
			s += nums[i]
			if s == k {
				count += 1
				continue
			}
		}
	}
	return count
}
