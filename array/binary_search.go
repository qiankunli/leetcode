package array

// 移动指针实现
func binarySearch3(nums []int, start, end, target int) int {

	l := start
	r := end

	for l <= r {
		mid := (start + end) / 2
		if nums[mid] == target {
			return mid
		}
		if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// 递归实现
func binarySearch2(nums []int, start, end, target int) int {
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}
	if target > nums[mid] && end > mid {
		return binarySearch(nums, mid+1, end, target)
	}
	if target < nums[mid] && mid > start {
		return binarySearch(nums, start, mid-1, target)
	}
	return -1
}
