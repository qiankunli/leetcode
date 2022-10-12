package array

// 153

// 需要始终将目标值（这里是最小值）套住，并不断收缩左边界或右边界。  始终将目标值套住，这句描述的好。

func findMin(nums []int) int {
	// l  r   如果大小是  l  < mid < r 说明没有旋转，最小值在最左边，可以收缩右边界
	// 如果 mid 小于 l 和 r，有旋转，最小值在左半边，可以收缩右边界
	// 如果中值 大于 l 和 r，有旋转，最小值在右半边，可以收缩左边界
	// l > md > r ，单调递减，不可能
	// 分析前面三种可能的情况，会发现情况1、2是一类，情况3是另一类。仅判断 mid 和r 的关系即可
	return -1
}

func findMin2(nums []int) int {

	//  一定要找到一个  点 比左右邻居都小，就是最小值，如果不是，就左右两边分别试下
	// 说白了，有点暴力法

	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	var result int
	var find func(nums []int, start, end int)
	find = func(nums []int, start, end int) {
		if start > end {
			return
		}
		l := start
		r := end

		mid := (l + r) / 2
		if nums[mid] < getElement(nums, mid+1) && nums[mid] < getElement(nums, mid-1) {
			result = nums[mid]
			return
		}

		find(nums, start, mid-1)
		find(nums, mid+1, end)
	}
	find(nums, 0, n-1)
	return result
}

func getElement(nums []int, index int) int {
	n := len(nums)
	if index < 0 {
		return nums[n+index]
	}
	if index >= n {
		return nums[index-n]
	}
	return nums[index]
}
