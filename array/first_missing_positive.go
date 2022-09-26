package array

import (
	"math"
	"sort"
)

// 41

func firstMissingPositive2(nums []int) int {
	// 求最小值（如果大于0 则为0） 最大值，for 最小到最大，map记录出现的数。空间复杂度是O(n)
	// max -min  等于n 则最后结果是max+1
	// 如果min > 1，则最后结果=1
	// 第三种方法，硬来
	n := len(nums)
	min := math.MaxInt
	max := math.MinInt
	for i := 0; i < n; i++ {
		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	if max-min >= n-1 {
		return max + 1
	}
	if min > 1 {
		return 1
	}
	// 一定有空隙，min 为0或负数
	for i := min + 1; i <= max-1; i++ {

	}
	return 0
}

// 答案要求O(n)，一次遍历是解决不了的，但是可以遍历多次
func firstMissingPositive(nums []int) int {
	sort.Ints(nums)

	cur := 1 // 当前关注的值
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}
		if nums[i] == cur {
			cur++
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		return cur
	}
	return cur
}
