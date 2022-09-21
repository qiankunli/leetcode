package array

import "math"

// 162
// 一个关键点是说，nums[i] < nums[i+1] 就往右走，一直递增的话，走到n-1，满足需求。不递增的话，也肯定能碰到一个。
// 题目要求logn ，但是想不透为啥，就有点尴尬
func findPeakElement(nums []int) int {

	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		return 0
	}
	l := 0
	r := n - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > get(nums, mid-1) && nums[mid] > get(nums, mid+1) {
			return mid
		} else if nums[mid] < get(nums, mid+1) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
func get(nums []int, index int) int {
	if index < 0 || index >= len(nums) {
		return math.MinInt
	}
	return nums[index]
}
