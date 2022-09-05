package array

import "sort"

func MaxFrequency(nums []int, k int) int {
	return maxFrequency(nums, k)
}

/*
1. 想到先对数据排序是不难的
*/

// 1838
func maxFrequency(nums []int, k int) int {
	if len(nums) == 1 {
		return 1
	}
	sort.Ints(nums)
	l := 0
	r := 1
	area := 0
	max := 0
	for i := 1; i < len(nums); i++ {
		area += (nums[i] - nums[i-1]) * (i - l)
		for area > k {
			area -= nums[i] - nums[l]
			l++
		}
		r = i
		if max < r-l+1 {
			max = r - l + 1
		}
	}
	return max
}
