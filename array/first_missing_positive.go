package array

import "sort"

// 41
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
