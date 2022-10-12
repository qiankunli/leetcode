package array

import (
	"math"
	"sort"
)

// 16
func threeSumClosest(nums []int, target int) int {
	// 暴力回溯，问题，部分解重复算。 一定是3个数
	// 先排序，然后固定第一个，左右指针流动第二个

	n := len(nums)

	if n < 3 {
		return sum(nums)
	}
	sort.Ints(nums)
	minAbs := math.MaxInt
	minSum := 0
	for i := 0; i < n-2; i++ {
		l := i + 1
		r := n - 1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if minAbs > abs(sum-target) {
				minAbs = abs(sum - target)
				minSum = sum
			}
			if sum == target {
				return minSum
			}
			if sum > target {
				r--
			} else {
				l++
			}
		}
	}
	return minSum
}
