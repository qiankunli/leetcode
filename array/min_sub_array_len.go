package array

import (
	"fmt"
	"math"
)

// 209
func minSubArrayLen(target int, nums []int) int {
	// 滑动窗口
	// 如果大于等于 target 右移

	n := len(nums)
	l := 0
	r := 0
	sum := 0
	min := math.MaxInt
	for r < n {
		sum += nums[r]
		for sum >= target && l <= n {
			if min < r-l+1 {
				min = r - l + 1
			}
			sum -= nums[l]
			l++
		}
		r++
	}
	return min
}

func minSubArrayLen2(target int, nums []int) int {
	// 不能排序   连续 窗口

	n := len(nums)
	l := 0
	r := 0

	// 问题，r 到底部循环 应该可以继续执行，
	sum := 0
	minLen := math.MaxInt
	for l <= r {
		if sum < target && r < n {
			sum += nums[r]
			r++
		}
		fmt.Println(sum, l, r)
		if sum >= target {
			if minLen > r-l {
				minLen = r - l
			}
			sum -= nums[l]
			l++
		}
	}
	return minLen
}
